package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/imattf/galere/controllers"
	"github.com/imattf/galere/templates"
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

	//Parse & Render home template
	tmpl := views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tmpl))

	//Parse & Render contact template
	tmpl = views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tmpl))

	//Parse & Render faq template
	tmpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.FAQ(tmpl))

	//Parse & Render goo template
	tmpl = views.Must(views.ParseFS(templates.FS, "goo.gohtml"))
	r.Get("/goo", controllers.StaticHandler(tmpl))

	r.NotFound(notfoundHandler)
	r.Get("/user/{userID}", userHandler)
	fmt.Println("Starting the galare server on :3000")
	http.ListenAndServe(":3000", r)
}
