package routes

import (
	"aplicacao-web/controllers"
	"net/http"
)

func Rotas() {
	http.HandleFunc("/", controllers.Index)
}
