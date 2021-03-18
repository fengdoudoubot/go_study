package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("index.html")
	// t.Execute(w, "Hello Template")

	t := template.Must(template.ParseFiles("index.html", "index2.html"))
	t.ExecuteTemplate(w, "index2.html", "我要去index2.html中")
}

func main() {

	http.HandleFunc("/testTemplate", testTemplate)

	http.ListenAndServe(":8080", nil)
}
