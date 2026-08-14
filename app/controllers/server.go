package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/kkato/todo-app/app/models"
	"github.com/kkato/todo-app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("GET /{$}", top)
	mux.HandleFunc("GET /signup", signupNew)
	mux.HandleFunc("POST /signup", signupCreate)
	mux.HandleFunc("GET /login", login)
	mux.HandleFunc("POST /authenticate", authenticate)
	mux.HandleFunc("GET /logout", logout)
	mux.HandleFunc("GET /todos", index)
	mux.HandleFunc("GET /todos/new", todoNew)
	mux.HandleFunc("POST /todos/save", todoSave)
	mux.HandleFunc("GET /todos/edit/{id}", todoEdit)
	mux.HandleFunc("POST /todos/update/{id}", todoUpdate)
	mux.HandleFunc("GET /todos/delete/{id}", todoDelete)

	return http.ListenAndServe(":"+config.Config.Port, mux)
}
