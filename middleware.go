package main

import (
	"fmt"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Testing Middleware :)")
		next.ServeHTTP(w,r)
	})
}