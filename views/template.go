package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
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
			"csrfField": func() template.HTML {
				return `<input type="hidden" />`
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

// func Parse(filepath string) (Template, error) {
// 	// parse the gohtml file
// 	tmpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		return Template{}, fmt.Errorf("parsing template: %w", err)
// 	}
// 	return Template{
// 		htmlTmpl: tmpl,
// 	}, nil
// }

type Template struct {
	htmlTmpl *template.Template
}

// helper function...
func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// render the gohtml file
	err := t.htmlTmpl.Execute(w, data)
	if err != nil {
		log.Printf("rendering error on %v", err)
		http.Error(w, "There was an error rendering the template.", http.StatusInternalServerError)
		return
	}
}
