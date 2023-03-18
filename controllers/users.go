package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/imattf/galere/context"
	"github.com/imattf/galere/models"
)

type Users struct {
	Templates struct {
		New            Template
		Signin         Template
		ForgotPassword Template
		CheckYourEmail Template
		ResetPassword  Template
	}
	UserService          *models.UserService
	SessionService       *models.SessionService
	PasswordResetService *models.PasswordResetService
	EmailService         *models.EmailService
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
	// cookie := http.Cookie{
	// 	Name:     "session",
	// 	Value:    session.Token,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// }
	// http.SetCookie(w, &cookie)
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
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
	// cookie := http.Cookie{
	// 	Name:     "session",
	// 	Value:    session.Token,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// }
	// http.SetCookie(w, &cookie)
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := context.User(ctx)
	// if user == nil {
	// 	http.Redirect(w, r, "/signin", http.StatusFound)
	// 	return
	// }
	fmt.Fprintf(w, "Current user: %s\n", user.Email)
}

func (u Users) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	err = u.SessionService.Delete(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	// Delete the user's cookie
	deleteCookie(w, CookieSession)
	http.Redirect(w, r, "/signin", http.StatusFound)
}

func (u Users) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.ForgotPassword.Execute(w, r, data)
	// u.Templates.CheckYourEmail.Execute(w, r, data)

}

func (u Users) ProcessForgotPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	pwReset, err := u.PasswordResetService.Create(data.Email)
	if err != nil {
		// TODO: Handle case where user/email doesn't exist
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	vals := url.Values{
		"token": {pwReset.Token},
	}
	// TODO: Make the URL configurable
	resetURL := "https://www.galere.com/reset-pw?" + vals.Encode()
	err = u.EmailService.ForgotPassword(data.Email, resetURL)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong.", http.StatusInternalServerError)
		return
	}
	// Don't render token here, user needs to validate from email
	u.Templates.CheckYourEmail.Execute(w, r, data)
}

func (u Users) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Token string
	}
	data.Token = r.FormValue("token")
	u.Templates.ResetPassword.Execute(w, r, data)
}

func (u Users) ProcessResetPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Token    string
		Password string
	}
	data.Token = r.FormValue("token")
	data.Token = r.FormValue("password")

	user, err := u.PasswordResetService.Consume(data.Token)
	if err != nil {
		fmt.Println(err)
		// TODO : Distinguish between server & invalid token errors
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	err = u.UserService.UpdatePassword(user.ID, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	// Sign the user in now that they have reset their password.
	// Any errors from this point onward should redirect to the sign in page.

	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

// User Middleware...
type UserMiddleware struct {
	SessionService *models.SessionService
}

func (umw UserMiddleware) SetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add logic for the SetUser middleware, then call next.ServeHTTP(w, r)

		// First try to read the cookie and set the user in the context if it can
		token, err := readCookie(r, CookieSession)
		if err != nil {
			// Cannot lookup the user with no cookie, so proceed w/out
			// user being set, then return
			next.ServeHTTP(w, r)
			return
		}

		// If we have a token, lookup the user by the token
		user, err := umw.SessionService.User(token)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		// If we get to here, we have a user that can store in the context

		// Get the context
		ctx := r.Context()

		// We need to derive a new context to store values.
		// Be certain to import our own context package (and not stdctx)
		ctx = context.WithUser(ctx, user)

		// Next we need to get a request that uses our new context.
		// We call WithContext function and it returns a new request w/ the context set.
		r = r.WithContext(ctx)

		// Finally we call the handler that our middleware was applied to
		// with the updated request.
		next.ServeHTTP(w, r)
	})
}

func (umw UserMiddleware) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := context.User(ctx)
		if user == nil {
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
