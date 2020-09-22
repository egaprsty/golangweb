package main

import (
	"log"
	"net/http"
	"v2/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/ga", handler.GaHandler)
	mux.HandleFunc("/ber", handler.BerHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)

	// Static Files

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("server started at localhost:3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

