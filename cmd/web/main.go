package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/wandirahman13/bookings/pkg/config"
	"github.com/wandirahman13/bookings/pkg/handlers"
	"github.com/wandirahman13/bookings/pkg/render"
)

// port
const portNumber = "3000"
const portApp = ":" + portNumber

// create var for app config
var app config.AppConfig

// session
var session *scs.SessionManager

// main application functionn
func main() {
	// set true for production
	app.InProduction = false

	// initialize new session using scs
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// send session to app config so it will be accessible for every part of app
	app.Session = session

	fmt.Println("=== Running Go Web App ===")
	fmt.Printf("Application is running on port %s . . .", portNumber)
	fmt.Println()

	// create template chache from main
	log.Println("[MAIN] create template cache from main")
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("[MAIN] cannot create template cache")
	}

	log.Println("[MAIN] assign template cache to app.TemplateCache in main.")
	app.TemplateCache = tc

	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	log.Println("[MAIN] trigger NewTemplates(&app) func from main")
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// _ = http.ListenAndServe(portApp, nil)

	srv := &http.Server{
		Addr:    portApp,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
