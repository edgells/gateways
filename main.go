package main

import (
	"fmt"
	"gateways/conf"
	"html"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var req int64

func init() {

	// load config

}

func main() {
	// 每个请求在 100ms 左右
	http.HandleFunc("*/",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "hello world%q", html.EscapeString(request.URL.Path))
			atomic.AddInt64(&req, 1)
			fmt.Printf("hello world %d\n", req)
			time.Sleep(time.Duration(100) * time.Millisecond)
		})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.HTTPPort), nil))
}
