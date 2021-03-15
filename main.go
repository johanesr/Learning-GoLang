package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/johanesr/jo_todo/config"
)

const port_number = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.IsProd = false

	// Set up session
	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.IdleTimeout = 30 * time.Minute
	session.Cookie.Persist = true //Cookie persist even after the browser is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProd

	app.Session = session
	// End session setup


	srv := &http.Server {
		Addr: port_number,
		Handler: routes(&app),
	}

	// Listen and Serve returns an error
	// log.fatal allows the app to return the error if there is any
	fmt.Println(fmt.Sprintf("Server is starting at port %s", port_number))
	log.Fatal(srv.ListenAndServe())
}
