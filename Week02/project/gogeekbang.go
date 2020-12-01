package main

import (
	"flag"
	"gogeekbang/internal/api"
	"gogeekbang/internal/dao"
	"gogeekbang/internal/pkg/config"
	"log"
	"net/http"
	"strconv"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "confPath", "app.ini", "配置文件路径")
	flag.Parse()

	config.Setup(confPath)
	dao.Setup(config.DatabaseConfig)
}

func main() {
	http.HandleFunc("/hello", api.Hello)
	http.HandleFunc("/getName", api.GetName)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.ServerConfig.HttpPort), nil))
}
