package main

import (
	"log"
	"net/http"

	"wsf/cours-devops/handler"
)

func main() {

	myHandler := handler.NewHandler()

	server := &http.Server{
		Addr:    ":8080",
		Handler: myHandler,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
