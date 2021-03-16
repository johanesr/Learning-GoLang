package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/johanesr/jo_todo/config"
	"github.com/johanesr/jo_todo/learnfunc"
)

func routes(app *config.AppConfig) http.Handler {
	// Router Handler
	mux := chi.NewRouter()

	//r.Use(WriteToConsole);
	//mux.Use(NoSurf) //CSRF validation do later
	mux.Use(SessionLoadSave)

	// Learn function routing
	mux.Get("/hello", learnfunc.PrintHello)
	mux.Get("/addition", learnfunc.Addition)
	mux.Get("/division", learnfunc.Division)
	mux.Get("/jsontest", learnfunc.JsonExample)
	mux.Get("/jsongettest", learnfunc.JsonGetExample)
	mux.Post("/jsonposttest", learnfunc.JsonPostExample)

	//Gorilla Mux
	//r := mux.NewRouter()
	//r.HandleFunc("/hello", learnfunc.PrintHello).Methods("GET")
	//r.HandleFunc("/addition", learnfunc.Addition).Methods("GET")
	//r.HandleFunc("/division", learnfunc.Division).Methods("GET")

	return mux
}