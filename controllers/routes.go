package controllers

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type helloController struct {
	templates helloTemplates
}

type helloTemplates struct {
	index *template.Template
}

func NewHelloController() controller {
	c := helloController{
		templates: helloTemplates{
			index: template.Must(template.ParseFiles("templates/index.html")),
		},
	}

	return func(r chi.Router) {
		r.Get("/hello", c.Hello)
	}
}

func (c *helloController) Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Default"
	}

	err := c.templates.index.Execute(w, map[string]string{"Name": name})
	if err != nil {
		panic(err)
	}
}
