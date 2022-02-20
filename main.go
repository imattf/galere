package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Windows OS proof approach...
	// parse the gohtml file
	// tplPath := filepath.Join("templates", "home.gohtml")
	// tpl, err := template.ParseFiles(tplPath)

	// unix/linux approach...
	// parse the gohtml file
	tpl, err := template.ParseFiles("templates/home.gohtml")

	if err != nil {
		log.Printf("parsing error on %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	// render the gohtml file
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("rendering error on %v", err)
		http.Error(w, "There was an error rendering the template.", http.StatusInternalServerError)
		return
	}

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:matthew@faulkners.io\">matthew@faulkners.io</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul><b>Is this thing free?</b> Yes, this is free.</ul>
	<ul><b>Would you like a wake-up call?</b> Yes, I need to quit my job & go back to college.</ul>
	<ul><b>Who can help me here?</b> Please email me at <a href="mailto:matthew@faulkners.io">matthew@faulkners.io.</a></ul>
	`)
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
