package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/imattf/galere/views"
)

// helper function...
// execute template
func executeTemplate(w http.ResponseWriter, filepath string) {
	// render the gohtml file
	tmpl, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing error on %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	// render the gohtml file
	tmpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// parse & render the gohtml file w/ the new helper function
	executeTemplate(w, "templates/home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	// parse & render the gohtml file w/ the new helper function
	executeTemplate(w, "templates/contact.gohtml")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	// parse & render the gohtml file w/ the new helper function
	executeTemplate(w, "templates/faq.gohtml")
}

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

	// enable chi logging across app
	// r.Use(middleware.Logger)

	// enable chi logging on single route
	// ... but doesn't work :(
	r.Route("/faqs", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/faq", faqHandler)
	})

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(notfoundHandler)
	r.Get("/user/{userID}", userHandler)
	fmt.Println("Starting the galare server on :3000")
	http.ListenAndServe(":3000", r)
}
