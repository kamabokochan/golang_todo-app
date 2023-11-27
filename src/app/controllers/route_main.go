package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	// 第二引数でデータを渡すことができる
	t.Execute(w, nil)
}
