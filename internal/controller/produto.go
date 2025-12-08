package controller

import (
	"API_GO/internal/model"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseGlob("assets/*.html"))
	temp.ExecuteTemplate(w, "index", model.BuscaTodosOsProdutos())
}
