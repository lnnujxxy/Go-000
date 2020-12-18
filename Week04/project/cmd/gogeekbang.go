package main

import (
	"context"
	"flag"
	"fmt"
	"gogeekbang/internal/api"
	"gogeekbang/internal/dao"
	"gogeekbang/internal/pkg/config"
	"gogeekbang/internal/pkg/db"
	"gogeekbang/internal/pkg/rdb"
	"gogeekbang/internal/service"
	"gogeekbang/internal/vars"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	conf := flag.String("conf", "conf/app.ini", "配置文件路径")
	flag.Parse()

	config.NewConfig(*conf)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// todo 接下来使用wire注入实现
	db := db.NewDB(config.DatabaseConfig)
	rdb := rdb.NewRedis(config.RedisConfig)
	dao := dao.NewDao(ctx, db, rdb)
	svc := service.NewService(ctx, dao)
	mux := http.NewServeMux()
	api.Route(mux, &svc)

	group, errCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		// 接收到信号量取消
		cancel()
		return nil
	})

	group.Go(func() error {
		server := &http.Server{
			Addr:           ":8000",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		go func() {
			if err := http.ListenAndServe(":" + strconv.Itoa(config.ServerConfig.HttpPort), wrapper(mux)); err != nil {
				// 异常取消
				cancel()
			}
		}()
		// 接收到取消后，执行关闭
		<-errCtx.Done()
		return server.Shutdown(ctx)
	})

	if err := group.Wait(); err != nil {
		log.Println("shutdown fail")
	}
}

func wrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := vars.NewVars(fmt.Sprintf("%v %v", r.Method, r.URL.String()))
		ctx := context.WithValue(context.Background(), "msg", v)
		r = r.WithContext(ctx)

		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic: %s", v.String())
			}
		}()

		h.ServeHTTP(w, r)
		log.Println(v.String())
	})
}
