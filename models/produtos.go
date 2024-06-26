package models

import "go-web-app/db"

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Id, Quantidade  int
}

func BuscaTodosOsProdutos() []Produto {
	database := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := database.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

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
	defer database.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	database := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := database.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer database.Close()
}

func DeletaProduto(id string) {
	database := db.ConectaComBancoDeDados()

	deletarProduto, err := database.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
	defer database.Close()
}

func EditaProduto(id string) Produto {
	database := db.ConectaComBancoDeDados()

	produtoBanco, err := database.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoAtualizar := Produto{}
	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoAtualizar.Id = id
		produtoAtualizar.Nome = nome
		produtoAtualizar.Descricao = descricao
		produtoAtualizar.Preco = preco
		produtoAtualizar.Quantidade = quantidade
	}
	defer database.Close()

	return produtoAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	database := db.ConectaComBancoDeDados()

	AtualizaProduto, err := database.Prepare("UPDATE produtos SET nome =$1, descricao=$2, preco=$3, quantidade=$4 WHERE id= $5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer database.Close()
}
