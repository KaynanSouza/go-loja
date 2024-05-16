package routes

import (
	"aplicacao-web/controllers"
	"net/http"
)

func Rotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/criar-produto", controllers.CreateProduct)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)

}
