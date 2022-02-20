package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// bio := `i've been hackin around 25 years`
	bio := `<script>alert("Haha, you have been h4x0r3d!");</script>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the Galare!</h1><p>User's bio: "+bio+"</p>")
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
