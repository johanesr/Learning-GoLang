package main

import (
	"github.com/gorilla/mux"
	"github.com/johanesr/jo_todo/config"
	"github.com/johanesr/jo_todo/learnfunc"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	// Router Handler
	r := mux.NewRouter()

	r.Use(WriteToConsole);

	// Learn function routing
	r.HandleFunc("/hello", learnfunc.PrintHello).Methods("GET")
	r.HandleFunc("/addition", learnfunc.Addition).Methods("GET")
	r.HandleFunc("/division", learnfunc.Division).Methods("GET")

	return r
}