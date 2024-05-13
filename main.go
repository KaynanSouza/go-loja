package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8080", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	produto := []Produto{
		{"Fone", "Preto", 59.2, 5},
		{"Cadeira", "Azul", 100, 15},
		{"Mouse", "Vermelho", 250, 3},
		{"NoteBook", "Branco", 1550.99, 170},
		{"Garrafa", "Roxo", 50.9, 50},
	}
	temp.ExecuteTemplate(writer, "Index", produto)
}
