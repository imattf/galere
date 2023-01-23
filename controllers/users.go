package controllers

import (
	"fmt"
	"net/http"

	"github.com/imattf/galere/models"
)

type Users struct {
	Templates struct {
		New    Template
		Signin Template
	}
	UserService    *models.UserService
	SessionService *models.SessionService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
		// CSRFField template.HTML
	}
	data.Email = r.FormValue("email")
	// data.CSRFField = csrf.TemplateField(r)
	u.Templates.New.Execute(w, r, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Email: ", r.FormValue("email"))
	// fmt.Fprint(w, "Password: ", r.FormValue("password"))
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		// TODO: Long term we shold show a warning about being
		// able to sign in as user
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/users/me", http.StatusFound)
	// fmt.Fprintf(w, "User created: %+v", user)
}

func (u Users) Signin(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Signin.Execute(w, r, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	// email, err := r.Cookie("email")
	// if err != nil {
	// 	fmt.Fprint(w, "the email cookie could not be read")
	// 	return
	// }
	// fmt.Fprintf(w, "Email cookie: %s\n", email.Value)
	// fmt.Fprintf(w, "Headers: %+v\n", r.Header)
	tokenCookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	user, err := u.SessionService.User(tokenCookie.Value)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Current user: %s\n", user.Email)
}
