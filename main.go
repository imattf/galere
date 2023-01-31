package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/imattf/galere/controllers"
	"github.com/imattf/galere/migrations"
	"github.com/imattf/galere/models"
	"github.com/imattf/galere/templates"
	"github.com/imattf/galere/views"
)

// Page Not Found...
func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Sorry... can't find that page!</h1>")
	fmt.Fprintln(w, "Page not found for ", r.URL.Path)
}

// Capture the chi URL Parameter...
func userHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the Galare</h1>")
	w.Write([]byte(fmt.Sprintf("hi %v", userID)))
}

func main() {

	// Setup Database...
	cfg := models.DefaultPostgresConfig()
	// fmt.Println(cfg.String()) //display connect string for Migration set-up
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//...migration tool goose
	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Setup Model Services...
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	// Setup Middleware...
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: Fix this before deploying to prod
		csrf.Secure(false),
	)

	// Setup Controllers...
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	usersC.Templates.Signin = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))

	// Setup Router and the Routes...
	r := chi.NewRouter()
	r.Use(csrfMw)      //apply csrf middleware functions
	r.Use(umw.SetUser) //apply user middile functions
	//...parse & render templates
	tmpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tmpl))
	tmpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tmpl))
	tmpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tmpl))
	tmpl = views.Must(views.ParseFS(templates.FS, "goo.gohtml", "tailwind.gohtml"))
	//...setup the routes
	r.Get("/goo", controllers.StaticHandler(tmpl))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.Signin)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	// r.Get("/users/me", usersC.CurrentUser)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
		// r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		// 	fmt.Fprint(w, "Hello")
		// })
	})
	r.Get("/user/{userID}", userHandler)
	r.NotFound(notfoundHandler)

	// Start the Server...
	fmt.Println("Starting the galare server on :3000")
	// http.ListenAndServe(":3000", csrfMw(umw.SetUser(r))) //with csrf Middleware wrapped in another Middleware
	http.ListenAndServe(":3000", r) // wrapped middleware applied above
}
