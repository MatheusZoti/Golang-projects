package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr: ":3000",
		Handler: router,

	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to Server", err)
	} 
}


func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}