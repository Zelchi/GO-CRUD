package controller

import (
	"API_GO/internal/web/model"
	"API_GO/internal/web/view"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseFS(view.ArquivosHTML, "*.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "home", produtos)
}

func NewItem(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Erro na conversão do preço: ", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			fmt.Println("Erro na conversão da quantidade: ", err)
		}

		model.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
