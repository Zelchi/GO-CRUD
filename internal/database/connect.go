package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	db, err := sql.Open("postgres", "user=admin dbname=admin password=admin host=localhost sslmode=disable")
	if err != nil {
		log.Fatal("Erro ao abrir conexão:", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Não foi possível conectar ao PostgreSQL:", err)
		return nil
	}

	sqlTabela := `
    CREATE TABLE IF NOT EXISTS produtos (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(255),
        descricao VARCHAR(255),
        preco DECIMAL(10,2),
        quantidade INTEGER
    )`

	if _, err := db.Exec(sqlTabela); err != nil {
		log.Fatal("Erro ao criar tabela:", err)
		return nil
	}

	return db
}
