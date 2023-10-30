package main

import (
	"log"
	"net/http"

	"github.com/eriicafes/twister-note/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// initialize middlewares
	r.Use(middleware.Logger)

	// apply controllers
	r.Group(controllers.NewHelloController())

	log.Fatal(http.ListenAndServe(":5000", r))
}
