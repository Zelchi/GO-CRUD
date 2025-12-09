package model

import (
	"API_GO/internal/database"
	"fmt"
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

	if len(produtos) == 0 {
		p := Produto{0, "Batata", "Batata", 0, 0}
		produtos = append(produtos, p)
	}

	fmt.Println(produtos)
	return produtos
}

func CriaNovoProduto(nome string, desc string, preco float64, quant int) {
	db := database.Connect()
	defer db.Close()

	state, err := db.Prepare(`
		INSERT INTO produtos(nome, descricao, preco, quantidade)
		VALUES($1, $2, $3, $4);
	`)
	if err != nil {
		log.Println("Erro ao preparar insert:", err)
		return
	}
	defer state.Close()

	if _, err := state.Exec(nome, desc, preco, quant); err != nil {
		log.Println("Erro ao inserir produto:", err)
	}
}
