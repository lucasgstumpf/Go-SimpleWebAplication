package controller

import (
	"aplicacaoWeb/models"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var tmp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.RecuperaTodosProdutos()
	tmp.ExecuteTemplate(w, "index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")

		precoInt, _ := strconv.Atoi(preco)

		quantidadeInt, _ := strconv.Atoi(quantidade)

		models.AdicionaProduto(nome, descricao, quantidadeInt, precoInt)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idResponse := r.URL.Query().Get("id")
	fmt.Printf("Requisição id = " + idResponse)
	models.RemoverProdutoById(idResponse)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entrou na funcao Edit\n")
	idProduto := r.URL.Query().Get("id")
	produto := models.RecuperaProdutoById(idProduto)
	tmp.ExecuteTemplate(w, "edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")
		precoInt, _ := strconv.Atoi(preco)
		quantidadeInt, _ := strconv.Atoi(quantidade)
		idInt, _ := strconv.Atoi(id)

		fmt.Println("Tudo lido com sucesso")

		models.AtualizaBancoById(idInt, nome, descricao, quantidadeInt, precoInt)

		http.Redirect(w, r, "/", 301)
	}
}
