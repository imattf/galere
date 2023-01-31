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

	//Parse & Render home template
	tmpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tmpl))

	//Parse & Render contact template
	tmpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tmpl))

	//Parse & Render faq template
	tmpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tmpl))

	//Parse & Render signup template
	// tmpl = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	// r.Get("/signup", controllers.StaticHandler(tmpl))

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()

	// Display connect string for Migration set-up
	// fmt.Println(cfg.String())

	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migration tool goose...
	// err = models.Migrate(db, "migrations")
	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	// Setup of controllers
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}

	// usersC := controllers.Users{}

	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))

	usersC.Templates.Signin = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.Signin)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Get("/users/me", usersC.CurrentUser)
	r.Get("/users/me", usersC.CurrentUser)
	// r.Get("/users/me", TimerMiddleware(usersC.CurrentUser))   //wrap with middleware

	//Parse & Render goo template
	tmpl = views.Must(views.ParseFS(templates.FS, "goo.gohtml", "tailwind.gohtml"))
	r.Get("/goo", controllers.StaticHandler(tmpl))

	r.NotFound(notfoundHandler)
	r.Get("/user/{userID}", userHandler)

	// Instantiate the middleware
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: Fix this before deploying to prod
		csrf.Secure(false),
	)

	fmt.Println("Starting the galare server on :3000")
	// http.ListenAndServe(":3000", TimerMiddleware(r.ServeHTTP))  //wrap with middleware
	// http.ListenAndServe(":3000", r)  //orig
	// http.ListenAndServe(":3000", csrfMw(r)) //with csrf Middleware
	http.ListenAndServe(":3000", csrfMw(umw.SetUser(r))) //with csrf Middleware wrapped in another Middleware

}

// Middleware function...
// func TimerMiddleware(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		h(w, r)
// 		fmt.Println("request time:", time.Since(start))
// 	}
// }
