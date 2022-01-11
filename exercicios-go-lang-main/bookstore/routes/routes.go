package routes

import (
	"net/http"

	"github.com/v1n1Fernand0/bookstore/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/update", controllers.Update)
}
