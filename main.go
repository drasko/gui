package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	staticDir = "static"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/hello.html")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
	}{"John Smith"}

	t.Execute(w, data)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/test.html")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
	}{"John Smith"}

	t.Execute(w, data)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve static files
	fs := http.FileServer(http.Dir(staticDir))
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})

	// Routes
	r.Get("/", rootHandler)
	r.Get("/test", testHandler)

	http.ListenAndServe(":3000", r)
}
