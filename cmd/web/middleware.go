package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

// just a rubbish function to learn creating a middleware
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[MIDDLEWARE] web got hit!")
		next.ServeHTTP(w, r)
	})
}

// middleware to create CSRF Token using nosurf to protect all POST requests
func CSRFTokenNoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad load and save session for every requests
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
