package model

import (
	"API_GO/database"
	"log"
	"net/http"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := database.Connect()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos ORDER BY id ASC")
	if err != nil {
		log.Println("Erro na query:", err)
		http.Error(nil, "Erro interno", 500)
		return nil
	}
	defer selectDeTodosOsProdutos.Close()

	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var p Produto
		err = selectDeTodosOsProdutos.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			log.Println("Erro no scan:", err)
			continue
		}
		produtos = append(produtos, p)
	}

	return produtos
}
