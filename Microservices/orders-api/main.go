package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	server := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(basicHandler),

	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to Server", err)
	} 
}


func basicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle GET
		if r.URL.Path == "/foo" {
			// Handle GET foo
		}
	}
	if r.Method == http.MethodPost {
		// Handle POST
	}
}