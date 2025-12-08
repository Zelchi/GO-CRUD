package main

import (
	"API_GO/internal/route"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Servidor rodando na porta 8000...")
	route.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
