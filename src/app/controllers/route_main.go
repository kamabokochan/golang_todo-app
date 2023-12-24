package controllers

import (
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := s.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodoByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}
