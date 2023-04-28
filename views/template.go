package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/imattf/galere/context"
	"github.com/imattf/galere/models"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tmpl := template.New(patterns[0])
	tmpl = tmpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField not implemented")
			},
			"currentUser": func() (*models.User, error) {
				return nil, fmt.Errorf("currentUser not implemented")
			},
			"errors": func() []string {
				// return []string{
				// 	"Don't do that!",
				// 	"The email address you provided is already associated with an account.",
				// 	"Something went wrong.",
				// }
				return nil
			},
		},
	)
	tmpl, err := tmpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTmpl: tmpl,
	}, nil
}

type Template struct {
	htmlTmpl *template.Template
}

// helper function...
func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error) {
	tmpl, err := t.htmlTmpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error rendering the page", http.StatusInternalServerError)
		return
	}

	tmpl = tmpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"currentUser": func() *models.User {
				return context.User(r.Context())
			},
			"errors": func() []string {
				var errorMessages []string
				for _, err := range errs {
					// TODO: Don't keep long term
					errorMessages = append(errorMessages, err.Error())
				}
				return errorMessages
			},
		},
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// render the gohtml file
	// err = tmpl.Execute(w, data)
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Printf("rendering error on %v", err)
		http.Error(w, "There was an error rendering the template.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
