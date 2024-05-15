package controllers

import (
	"aplicacao-web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscarProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "criar-produto", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro ao converter o preço", err)
			return
		}

		quantidadeConv, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro ao converter a quantidade", err)
			return
		}

		models.CreateProdict(nome, descricao, precoConv, quantidadeConv)
	}

	http.Redirect(w, r, "/", 301)
}
