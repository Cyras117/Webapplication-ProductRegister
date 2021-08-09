package controllers

import (
	m "go_cadastro/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := m.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

//delete?id={{.Id}}
func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	m.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Panicln("Erro na conversão do preço")
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Panicln("Erro na conversão da quantidade")
		}

		m.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)

}
