package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var db *sql.DB
var temp = template.Must(template.ParseGlob("assets/*.html"))

func main() {
	conexao := "user=admin dbname=admin password=admin host=localhost sslmode=disable"
	var err error
	db, err = sql.Open("postgres", conexao)
	if err != nil {
		log.Fatal("Erro ao abrir conexão:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Não foi possível conectar ao PostgreSQL:", err)
	}

	sqlTabela := `
    CREATE TABLE IF NOT EXISTS produtos (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(255),
        descricao VARCHAR(255),
        preco DECIMAL(10,2),
        quantidade INTEGER
    )`

	_, err = db.Exec(sqlTabela)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}

	fmt.Println("Servidor rodando na porta 8000...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	selectDeTodosOsProdutos, err := db.Query("select * from produtos ORDER BY id ASC")
	if err != nil {
		log.Println("Erro na query:", err)
		http.Error(w, "Erro interno", 500)
		return
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

	temp.ExecuteTemplate(w, "index", produtos)
}
