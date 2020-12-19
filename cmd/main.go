package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	staticDir = "gui/static"
)

var (
	t *template.Template
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Name string
	}{"John Smith"}

	t.ExecuteTemplate(w, "index", data)
}

func thingsHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
	}{"John Smith"}

	t.ExecuteTemplate(w, "things", data)
}

func channelsHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
	}{"John Smith"}

	t.ExecuteTemplate(w, "channels", data)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve static files
	fs := http.FileServer(http.Dir(staticDir))
	r.Handle("/*", fs)

	// Routes
	r.Get("/", indexHandler)
	r.Get("/things", thingsHandler)
	r.Get("/channels", channelsHandler)

	var err error
	t, err = template.ParseGlob("gui/views/*")
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":3000", r)
}
