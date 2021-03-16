package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/johanesr/jo_todo/config"
	_ "github.com/go-sql-driver/mysql"
)

const port_number = ":8080"
var app config.AppConfig
var session *scs.SessionManager
var sqlDB *sql.DB

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

	fmt.Println("Connecting to a MySQL Databse")
	var err error
	sqlDB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/student")
	//db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/student")
	if err!=nil {log.Fatal("An error when connecting to mysql database: ",err)}
	defer db.Close()
	fmt.Println("Successfully connected to a MySQL Database!")

	Insertion(sqlDB, "test3", "Sport Science", 2)
	Selection(sqlDB)

	srv := &http.Server {
		Addr: port_number,
		Handler: routes(&app),
	}

	// Listen and Serve returns an error
	// log.fatal allows the app to return the error if there is any
	fmt.Println(fmt.Sprintf("Server is starting at port %s", port_number))
	log.Fatal(srv.ListenAndServe())
}
