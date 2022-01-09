package main

import (
	"net/http"

	"github.com/v1n1Fernand0/bookstore/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
