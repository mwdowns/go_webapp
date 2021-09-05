package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf middleware function to use no_surf mod to prevent cross-site-request-forgery attacks
func NoSurf(next http.Handler) http.Handler {
	csrf_handler := nosurf.New(next)

	csrf_handler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrf_handler
}

// SessionLoad loads and saves session on request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
