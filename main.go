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

type testDB struct {
	Id 		int
	Name 	sql.NullString
}

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
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/student")
	if err!=nil {log.Fatal("An error when connecting to mysql database: ",err)}
	defer db.Close()
	fmt.Println("Successfully connected to a MySQL Database!")

	//insert, err := db.Query("INSERT INTO student VALUES (6,'go', 'Computer Science', 1)")
	//if err!=nil {log.Fatal("An error when inserting to mysql database: ",err)}
	//defer insert.Close()
	//fmt.Println("Successfully inserted to a MySQL Database!")

	results, err := db.Query("SELECT student_id, student_name FROM student")
	if err!=nil {log.Fatal("An error when reading from mysql database: ",err)}
	for results.Next() {
		var user testDB

		// Scan have to pass a number of arguments = number of column return
		err := results.Scan(&user.Id, &user.Name)
		if err!=nil {log.Print("fatal: ",err)}
		fmt.Printf("ID: %d, name: %s\n",user.Id,user.Name.String)
	}

	srv := &http.Server {
		Addr: port_number,
		Handler: routes(&app),
	}

	// Listen and Serve returns an error
	// log.fatal allows the app to return the error if there is any
	fmt.Println(fmt.Sprintf("Server is starting at port %s", port_number))
	log.Fatal(srv.ListenAndServe())
}
