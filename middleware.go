package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// Info from gorilla/mux:
// Call the next handler, which can be another middleware in the chain, or the final handler.
// in this case next is r handler from routes.go which is also the final handler
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Testing Middleware :)")
		next.ServeHTTP(w,r)
	})
}

// Middleware to check CSRF Token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next);

	// To learn more about HTTP cookies
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.IsProd, // HTTPS (true/false) => Development or Production
		SameSite: http.SameSiteLaxMode,
		// Lax allows cookies to be sent if user comes from external link
	})

	return csrfHandler;
}

// Loads and Save sessions on every request
func SessionLoadSave(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}