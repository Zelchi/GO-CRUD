package controller

import (
	"API_GO/internal/web/model"
	"API_GO/internal/web/view"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseFS(view.ArquivosHTML, "*.html"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "home", produtos)
}

func AddItemPage(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := model.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "edit", produto)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		model.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	id := r.FormValue("id")
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erro na conversão do ID para int:", err)
		return
	}
	floatPreco, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("Erro na conversão do Preço para float64:", err)
		return
	}
	intQuantidade, err := strconv.Atoi(quantidade)
	if err != nil {
		log.Println("Erro na conversão da Quantidade para int:", err)
		return
	}

	model.AtualizaProduto(intId, nome, descricao, floatPreco, intQuantidade)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")

	fmt.Println("hit")

	idConvertido, err := strconv.Atoi(idDoProduto)
	if err != nil {
		log.Println("Erro na conversão do ID do produto: ", err)
	} else {
		model.DeletaProduto(idConvertido)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
