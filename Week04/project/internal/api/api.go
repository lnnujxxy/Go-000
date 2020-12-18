package api

import (
	"fmt"
	"gogeekbang/internal/service"
	"net/http"
)

func HelloHandler(_ *service.Service) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n")
	}
}

func GetNameHandler(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name, err := svc.GetName(1)
		if err != nil {
			// todo error code
		}
		fmt.Fprintf(w, "name:%s", name)
	}
}
