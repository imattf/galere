package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/imattf/galere/controllers"
	"github.com/imattf/galere/views"
)

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Sorry... can't find that page!</h1>")
	fmt.Fprintln(w, "Page not found for ", r.URL.Path)
}

// Capture the chi URL Parameter
func userHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the Galare</h1>")
	w.Write([]byte(fmt.Sprintf("hi %v", userID)))
}

func main() {

	// a new chi router
	r := chi.NewRouter()

	//Parse home template
	tmpl, err := views.Parse("templates/home.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tmpl))

	//Parse contact template
	tmpl, err = views.Parse("templates/contact.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tmpl))

	//Parse faq template
	tmpl, err = views.Parse("templates/faq.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tmpl))

	r.NotFound(notfoundHandler)
	r.Get("/user/{userID}", userHandler)
	fmt.Println("Starting the galare server on :3000")
	http.ListenAndServe(":3000", r)
}
