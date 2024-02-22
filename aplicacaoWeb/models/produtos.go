package models

import (
	"aplicacaoWeb/db"
	"fmt"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      int
	Quantidade int
}

func AdicionaProduto(nome string, descricao string, quantidade int, preco int) {
	db := db.ConectaBanco()

	insertProdutosBanco, err := db.Prepare("INSERT INTO produtos(nome,descricao,quantidade,valor) values ($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insertProdutosBanco.Exec(nome, descricao, quantidade, preco)

	defer db.Close()

}

func RecuperaTodosProdutos() []Produto {

	produtos := []Produto{}

	db := db.ConectaBanco()
	selectResposta, _ := db.Query("SELECT * FROM produtos")

	for selectResposta.Next() {
		var id, preco, quantidade int
		var nome, descricao string

		err := selectResposta.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			fmt.Printf("Deu ruim")
		}

		p := Produto{id, nome, descricao, preco, quantidade}

		produtos = append(produtos, p)

	}

	defer db.Close()
	return produtos
}

func RecuperaProdutoById(idProduto string) Produto {
	db := db.ConectaBanco()

	responseBusca, _ := db.Query("select * from produtos where id=$1", idProduto)

	var p Produto

	for responseBusca.Next() {
		var id, quantidade, preco int
		var nome, descricao string

		responseBusca.Scan(&id, &nome, &descricao, &preco, &quantidade)

		p = Produto{id, nome, descricao, preco, quantidade}

	}

	return p
}

func RemoverProdutoById(idProduto string) {
	db := db.ConectaBanco()

	queryDelete, err := db.Prepare("DELETE FROM produtos where id=$1")
	if err != nil {
		fmt.Println("Erro na query")
		panic(err.Error())
	}

	queryDelete.Exec(idProduto)

	defer db.Close()

}

func AtualizaBancoById(id int, nome string, descricao string, quantidade int, preco int) {
	db := db.ConectaBanco()

	queryUpdate, err := db.Prepare("update produtos set nome=$1, descricao=$2, valor=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
		fmt.Println("\n\n\n\n")
	}

	queryUpdate.Exec(nome, descricao, quantidade, preco, id)

	defer db.Close()

}
