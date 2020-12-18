package api

import (
	"fmt"
	"gogeekbang/internal/service"
	"net/http"
	"strconv"
)

func HelloHandler(_ *service.Service) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n")
	}
}

func GetNameHandler(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		name, err := svc.GetName(id)
		if err != nil {
			// todo error code
		}
		// todo wrapper response
		fmt.Fprintf(w, "name:%s, id:%d", name, id)
	}
}
