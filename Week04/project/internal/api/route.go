package api

import (
	"gogeekbang/internal/service"
	"net/http"
)

func Route(mux *http.ServeMux, service *service.Service) {
	mux.HandleFunc("/hello", HelloHandler(service))
	mux.HandleFunc("/getName", GetNameHandler(service))
}