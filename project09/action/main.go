package main

import (
	"html/template"
	"net/http"
	"go_code/project09/action/model"
)

func testRange(w http.ResponseWriter, r *http.Request) {
	
	t := template.Must(template.ParseFiles("range.html"))
	var emps []*model.Employee
	emp := &model.Employee{
		ID:       1,
		LastName: "lll",
		Email:    "lxl@qq.com",
	}
	emps = append(emps, emp)
	emp2 := &model.Employee{
		ID:       2,
		LastName: "vvv",
		Email:    "bbh@sina.com",
	}
	emps = append(emps, emp2)
	emp3 := &model.Employee{
		ID:       3,
		LastName: "mmm",
		Email:    "mm@sina.com",
	}
	emps = append(emps, emp3)

	t.Execute(w, emps)
}


func testIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("if.html"))
	age := 17
	t.Execute(w, age > 18)
}

func testWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("with.html"))
	t.Execute(w, "狸猫")
}


func testTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template1.html", "template2.html"))
	t.Execute(w, "我能在两个文件中显示吗？")
}

func testDefine(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("define.html"))
	t.ExecuteTemplate(w, "model", "")
}

func testDefine2(w http.ResponseWriter, r *http.Request) {
	age := 17
	var t *template.Template
	if age < 18 {
		t = template.Must(template.ParseFiles("define2.html"))
	} else {
		t = template.Must(template.ParseFiles("define2.html", "content1.html"))
	}
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/testIf", testIf)
	http.HandleFunc("/testRange", testRange)
	http.HandleFunc("/testWith", testWith)
	http.HandleFunc("/testTemplate", testTemplate)
	http.HandleFunc("/testDefine", testDefine)
	http.HandleFunc("/testDefine2", testDefine2)

	http.ListenAndServe(":8080", nil)
}
