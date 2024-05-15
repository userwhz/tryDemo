package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./go/gin/template/index.tmpl")
	if err != nil {
		fmt.Printf("template parse failed, err: %v\n", err)
	}
	//渲染模板
	t.Execute(w, "ggggg")
}

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}
}
