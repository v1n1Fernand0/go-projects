package models

import "github.com/v1n1Fernand0/bookstore/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func ListarProdutos() []Produto {

	db := db.ConectaBanco()
	listaProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for listaProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = listaProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func ListarPorId(id int) Produto {
	db := db.ConectaBanco()
	query, erro := db.Query("select * from produtos where id=$1", id)
	if erro != nil {
		panic(erro.Error())
	}
	p := Produto{}
	for query.Next() {
		query.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
	}

	defer db.Close()
	return p
}

func CriarProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()
	query, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	query.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func Delete(id string) {
	db := db.ConectaBanco()
	delete, erro := db.Prepare("delete from produtos where id=$1")
	if erro != nil {
		panic(erro.Error())
	}

	delete.Exec(id)
	db.Close()
}

func Atualizar(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()
	query, erro := db.Prepare("update produtos set nome =$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if erro != nil {
		panic(erro.Error())
	}
	query.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
