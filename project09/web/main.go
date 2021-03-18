package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}


func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w , "hello!",r.URL.Path)
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "通过自己创建的处理器处理请求！")
	fmt.Fprintln(w, "通过详细配置服务器的信息来处理请求！")
}

func main()  {
	// http.HandleFunc("/",handler)
	// http.ListenAndServe(":8080",nil)

	myHandler := MyHandler{}
	// http.Handle("/myHandler", &myHandler)		
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}

	server.ListenAndServe()
}