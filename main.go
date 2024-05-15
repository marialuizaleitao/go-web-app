package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Camiseta preta de Kung Fu", 29.90, 10},
		{"Boné", "Bóne vermelho Aba Reta", 25.99, 2},
		{"Pulseira", "Pulseira prata com detalhes verdes", 159.99, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
