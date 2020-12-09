package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		// 接收到信号量取消
		cancel()
	}()

	group, errCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		server := &http.Server{
			Addr:           ":8080",
			Handler:        &helloHandler{},
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		go func() {
			if err := server.ListenAndServe(); err != nil {
				// 异常取消
				cancel()
			}
		}()
		// 接收到取消后，执行关闭
		<-errCtx.Done()
		fmt.Println("server 1 exit")
		return server.Shutdown(ctx)
	})

	group.Go(func() error {
		server := &http.Server{
			Addr:           ":8081",
			Handler:        &helloHandler{},
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		go func() {
			if err := server.ListenAndServe(); err != nil {
				cancel()
			}
		}()

		<-errCtx.Done()
		fmt.Println("server 2 exit")
		return server.Shutdown(ctx)
	})

	if err := group.Wait(); err != nil {
		log.Println("shutdown fail")
	}
}

type helloHandler struct {
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
