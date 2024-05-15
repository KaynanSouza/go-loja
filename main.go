package main

import (
	"aplicacao-web/routes"
	"net/http"
)

func main() {
	routes.Rotas()
	http.ListenAndServe(":8080", nil)
}
