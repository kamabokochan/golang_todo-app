package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
	"todo_app/src/app/models"
	"todo_app/src/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("src/app/views/templates/%s.html", file))
	}

	// MEMO
	// Must内ParseFilesがエラーの場合は、panicになる
	// mainを実行する前に落とすことで安全に初期化する際に利用されるらしいが
	// 今回はmain内で実行されているので、今回の利用目的は要調査。（初期化処理という意味合いで利用している？）一旦講座通りに実装を進める
	templates := template.Must(template.ParseFiles(files...))

	// MEMO
	// 指定されたデータオブジェクトに関連付けられているテンプレートを適用し、出力をwrに書き込む
	// layoutが受け口になるため、明示的にlayoutと指定、dataをlayoutを通し各テンプレートへ読み込ませていると一旦理解
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (s models.Session, err error) {
	cookie, err := r.Cookie("myapp_uuid")
	if err != nil {
		return s, err
	}

	s = models.Session{UUID: cookie.Value}
	if ok, _ := s.CheckSession(); !ok {
		err = fmt.Errorf("Invalid session")
	}
	return s, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

func StartMainServer() error {
	// 公開したいディレクトリを指定
	files := http.FileServer(http.Dir(config.Config.Static))
	// http.StripPrefix()は、第1引数に指定したパスを、http.FileServer()が捜索するURLから取り除く
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit)) // handler関数に引数idを渡すため、edit/と最後に/をつけている
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
