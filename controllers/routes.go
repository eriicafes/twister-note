package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type helloController struct{}

func NewHelloController() controller {
	c := helloController{}

	return func(r chi.Router) {
		r.Get("/hello", c.Hello)
	}
}

func (c *helloController) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<div>hello world my name is <b>Luq</b></div>")
}
