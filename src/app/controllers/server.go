package controllers

import (
	"net/http"
	"todo_app/src/config"
)

func StartMainServer() error {
	// 公開したいディレクトリを指定
	files := http.FileServer(http.Dir(config.Config.Static))
	// http.StripPrefix()は、第1引数に指定したパスを、http.FileServer()が捜索するURLから取り除く
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
