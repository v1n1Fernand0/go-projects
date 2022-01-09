package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/v1n1Fernand0/bookstore/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "index", models.ListarProdutos())
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		precoFloat, erroPreco := strconv.ParseFloat(preco, 64)
		quantidadeInt, erroQuantidade := strconv.Atoi(quantidade)
		RetornaErroLog(erroPreco, "preco")
		RetornaErroLog(erroQuantidade, "quantidade")

		models.CriarProduto(nome, descricao, precoFloat, quantidadeInt)
		http.Redirect(w, r, "/", 301)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		idInt, erroIntId := strconv.Atoi(id)
		precoFloat, erroFloatPreco := strconv.ParseFloat(preco, 64)
		quantidadeInt, erroIntQuantidade := strconv.Atoi(quantidade)

		RetornaErroLog(erroIntId, "id")
		RetornaErroLog(erroFloatPreco, "preco")
		RetornaErroLog(erroIntQuantidade, "quantidade")
		models.Atualizar(idInt, nome, descricao, precoFloat, quantidadeInt)
		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.Delete(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idProduto, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}
	produto := models.ListarPorId(idProduto)

	temp.ExecuteTemplate(w, "edit", produto)
}

func RetornaErroLog(erro error, nomeProp string) {
	if erro != nil {
		log.Println("Erro na propriedade ", nomeProp, "convertida!")
		panic(erro.Error())
	}
}
