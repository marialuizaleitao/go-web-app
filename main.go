package main

import (
	"go-web-app/db"
	"html/template"
	"log"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Id, Quantidade  int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// usa um pool de conexões para reutilizar conexões existentes
	database := db.PoolDeConexoes()

	selectDeTodosOsProutos, err := database.Query("SELECT nome, descricao, preco, quantidade FROM produtos")
	if err != nil {
		log.Fatal(err)
	}
	defer selectDeTodosOsProutos.Close()

	var produtos []Produto

	for selectDeTodosOsProutos.Next() {
		var produto Produto

		err := selectDeTodosOsProutos.Scan(&produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			log.Fatal(err)
		}
		produtos = append(produtos, produto)
	}

	// verifica se houve algum erro durante a iteração
	if err := selectDeTodosOsProutos.Err(); err != nil {
		log.Fatal(err)
	}

	// executa o template com os produtos
	temp.ExecuteTemplate(w, "Index", produtos)
}
