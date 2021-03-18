package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
		MaxAge:   60,
	}
	cookie2 := http.Cookie{
		Name:     "user2",
		Value:    "admin2",
		HttpOnly: true,
	}
	// w.Header().Set("Set-Cookie", cookie.String())
	// w.Header().Add("Set-Cookie", cookie2.String())
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}


func getCookies(w http.ResponseWriter, r *http.Request) {
	//获取请求头中所有的Cookie
	// cookies := r.Header["Cookie"]
	
	cookie, _ := r.Cookie("user")
	fmt.Println("得到的Cookie有：", cookie)
}

func main() {

	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookies", getCookies)

	http.ListenAndServe(":8080", nil)
}
