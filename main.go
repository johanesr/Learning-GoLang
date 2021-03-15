package main

import (
	"fmt"
	"github.com/johanesr/jo_todo/config"
	"log"
	"net/http"
)

const port_number = ":8080"

func main() {
	var app config.AppConfig

	srv := &http.Server {
		Addr: port_number,
		Handler: routes(&app),
	}

	fmt.Println(fmt.Sprintf("Server is starting at port %s", port_number))
	log.Fatal(srv.ListenAndServe())


	// Listen and Serve returns an error
	// log.fatal allows the app to return the error if there is any
}
