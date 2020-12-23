package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my app!")
}

func welcomeV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to v1!")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func main() {
	log.Printf("starting applicaiton...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(func(r http.Handler) http.Handler {
		log.Printf("started application")
		return r
	})
	r.Get("/", welcome)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/", welcomeV1)
		r.Get("/helloWorld", helloWorld)
	})
	http.ListenAndServe(":8000", r)
	log.Printf("shutdown applicaiton...")
}
