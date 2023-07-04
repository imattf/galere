package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/imattf/galere/context"
	"github.com/imattf/galere/errors"
	"github.com/imattf/galere/models"
)

type Galleries struct {
	Templates struct {
		Show  Template
		New   Template
		Edit  Template
		Index Template
	}
	GalleryService *models.GalleryService
}

func (g Galleries) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Title string
	}
	data.Title = r.FormValue("title")
	g.Templates.New.Execute(w, r, data)
}

func (g Galleries) Create(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserID int
		Title  string
	}
	data.UserID = context.User(r.Context()).ID
	data.Title = r.FormValue("title")

	gallery, err := g.GalleryService.Create(data.Title, data.UserID)
	if err != nil {
		g.Templates.New.Execute(w, r, data, err)
		return
	}
	editPath := fmt.Sprintf("galleries/%d/edit", gallery.ID)
	http.Redirect(w, r, editPath, http.StatusFound)
	// g.Templates.New.Execute(w, r, data)
}

func (g Galleries) Edit(w http.ResponseWriter, r *http.Request) {
	//Get the ID of the gallery we are editing
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// 404 error - page isn't found
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}
	//Query for the gallery
	gallery, err := g.GalleryService.ByID(id)
	if err != nil {
		// if err == models.ErrNotFound {
		if errors.Is(err, models.ErrNotFound) {
			http.Error(w, "Gallery not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	//Verify user owns the gallery
	user := context.User((r.Context()))
	if gallery.UserID != user.ID {
		http.Error(w, "You are not authorized to edit this gallery", http.StatusForbidden)
		return
	}
	//Render and Edit the gallery
	// data := struct {
	// 	ID    int
	// 	Title string
	// }{
	// 	ID:    gallery.ID,
	// 	Title: gallery.Title,
	// }
	var data struct {
		ID    int
		Title string
	}
	data.ID = gallery.ID
	data.Title = gallery.Title
	g.Templates.Edit.Execute(w, r, data)
}

func (g Galleries) Update(w http.ResponseWriter, r *http.Request) {
	//Get the ID of the gallery we are editing
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// 404 error - page isn't found
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}
	//Query for the gallery
	gallery, err := g.GalleryService.ByID(id)
	if err != nil {
		// if err == models.ErrNotFound {
		if errors.Is(err, models.ErrNotFound) {
			http.Error(w, "Gallery not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	//Verify user owns the gallery
	user := context.User((r.Context()))
	if gallery.UserID != user.ID {
		http.Error(w, "You are not authorized to edit this gallery", http.StatusForbidden)
		return
	}

	//Parse the title from the form and update the gallery
	title := r.FormValue("title")
	gallery.Title = title
	err = g.GalleryService.Update(gallery)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	editPath := fmt.Sprintf("/galleries/%d/edit", gallery.ID)
	http.Redirect(w, r, editPath, http.StatusFound)
}

func (g Galleries) Index(w http.ResponseWriter, r *http.Request) {
	type Gallery struct {
		ID    int
		Title string
	}
	var data struct {
		Galleries []Gallery
	}
	//Verify user owns the gallery
	user := context.User(r.Context())
	galleries, err := g.GalleryService.ByUserID(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	for _, gallery := range galleries {
		data.Galleries = append(data.Galleries, Gallery{
			ID:    gallery.ID,
			Title: gallery.Title,
		})
	}
	// TODO: Lookup the galleries we are going to render
	g.Templates.Index.Execute(w, r, data)
}

func (g Galleries) Show(w http.ResponseWriter, r *http.Request) {
	//Get the ID of the gallery we are editing
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// 404 error - page isn't found
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}
	//Query for the gallery
	gallery, err := g.GalleryService.ByID(id)
	if err != nil {
		// if err == models.ErrNotFound {
		if errors.Is(err, models.ErrNotFound) {
			http.Error(w, "Gallery not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	var data struct {
		ID     int
		Title  string
		Images []string
	}
	data.ID = gallery.ID
	data.Title = gallery.Title

	// Load some temporary images
	for i := 0; i < 20; i++ {
		w, h := rand.Intn(500)+200, rand.Intn(500)+200
		catImageURL := fmt.Sprintf("https://placekitten.com/%d/%d", w, h)
		data.Images = append(data.Images, catImageURL)
	}
	g.Templates.Show.Execute(w, r, data)
}
