package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tmpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func FAQ(tmpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is this thing free?",
			Answer:   "Yes, this is free.",
		},
		{
			Question: "Would you like a wake-up call?",
			Answer:   "Yes, I need to quit my job & go back to college.",
		},
		{
			Question: "Who can help me here?",
			Answer:   `Please email me at <a href="mailto:matthew@faulkners.io">matthew@faulkners.io.</a>`,
		},
		{
			Question: "Where is your office located?",
			Answer:   "Over there...",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, questions)
	}
}
