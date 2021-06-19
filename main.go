package main

import (
	// my api
	"knn-golang/api"
	"log"
	// Package http provides HTTP client and server implementations. Get, Head, Post, and PostForm make HTTP (or HTTPS) requests:
	"net/http"
	// Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.
	"github.com/gorilla/mux"
	// Allow * With Credentials Security Protection
	"github.com/rs/cors"
)

func main() {

	//initialize router
	r := mux.NewRouter()
	//call predict method for this path
	r.HandleFunc("/predict", api.Predict).Methods("POST")

	//allow cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	exit := make(chan error)

	//start a goroutine for the backend server
	go func() {
		log.Fatal(http.ListenAndServe(":3000", handler))
		exit <- nil

	}()
	serveUI := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	serveUI.Handle("/", fs)
	//start a goroutine for the front end files
	go func() {
		log.Println(http.ListenAndServe(":3001", serveUI))
	}()

	<-exit
}
