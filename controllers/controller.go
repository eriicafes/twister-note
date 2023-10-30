package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type controller func(chi.Router)

func (c controller) Router() chi.Router {
	r := chi.NewRouter()
	c(r)
	return r
}

func (c controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Router().ServeHTTP(w, r)
}
