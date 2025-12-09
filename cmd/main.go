package main

import (
	"API_GO/internal/route"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Servidor rodando na porta 8000...")
	route.LoadBack()
	route.LoadFront()
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Não foi possivel iniciar o servidor na porta 8000")
	}
}
