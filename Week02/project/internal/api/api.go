package api

import (
	"fmt"
	"gogeekbang/internal/service"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func Hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func GetName(w http.ResponseWriter, r *http.Request) {
	name, err := service.GetName(2)
	// 在api层统一error处理
	if err != nil {
		log.Print("api: GetName err=", errors.Cause(err).Error())
	}

	fmt.Fprintf(w, "name:%s", name)
}
