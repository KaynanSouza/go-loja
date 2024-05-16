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
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro ao converter o preço", err)
		}

		quantidadeConv, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro ao converter a quantidade", err)
		}

		models.CreateProdict(nome, descricao, precoConv, quantidadeConv)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeleteProduct(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditProduct(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro ao converter o Id para int: ", err)
		}

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter o preco em float: ", err)
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter a quantidade em int: ", err)
		}

		models.UpdateProduct(idConv, nome, descricao, precoConv, quantidadeConv)

	}

	http.Redirect(w, r, "/", 301)
}
