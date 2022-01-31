package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the Galare!</h1>")
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

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		// TODO: first one
// 		// http.Error(w, "Page not found", http.StatusNotFound)
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintln(w, "Page not found for ", r.URL.Path)
// 		fmt.Fprintln(w, "Page not found for ", r.URL.RawPath)
// 	}

// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}

}

func main() {
	var router Router
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the galare server on :3000")
	http.ListenAndServe(":3000", router)
}
