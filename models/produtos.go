package models

import (
	"go_cadastro/dbcon"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func DeletaProduto() {
	db := dbcon.ConectaBD()
	deletarProduto := db.Prepare("delete from produtos ")
	defer db.Close()
}

func BuscaTodosOsProdutos() []Produto {
	db := dbcon.ConectaBD() //conectaBD()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quantidade int

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
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

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := dbcon.ConectaBD()
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao,preco,quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
