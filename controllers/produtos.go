package controllers

import (
	"aplicacao-web/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscarProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
