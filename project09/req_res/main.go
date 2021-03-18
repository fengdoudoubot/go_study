package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"go_code/project09/model"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "请求地址是：", r.URL.Path)
	// fmt.Fprintln(w, "请求地址后的查询字符串是：", r.URL.RawQuery)
	// fmt.Fprintln(w, "请求头：", r.Header)
	// fmt.Fprintln(w, "Accept-Encoding：", r.Header["Accept-Encoding"])
	// fmt.Fprintln(w, "Accept-Encoding的属性值是：", r.Header.Get("Accept-Encoding"))
	
	// len := r.ContentLength
	// body := make([]byte, len)
	// r.Body.Read(body)
	// fmt.Fprintln(w, "请求体中的内容有：", string(body))

	// r.ParseForm()
	// fmt.Fprintln(w, "请求参数有：", r.Form)
	// fmt.Fprintln(w, "POST请求的form表单中的请求参数有：", r.PostForm)
	fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("user"))
	fmt.Fprintln(w, "Form表单中的username请求参数的值是：", r.PostFormValue("username"))
}

func testJsonRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := model.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Email:    "admin@qq.com",
	}
	json, _ := json.Marshal(user)
	w.Write(json)
}

func testRedire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testJsonRes)
	http.HandleFunc("/testRedirect", testRedire)

	http.ListenAndServe(":8080", nil)
}
