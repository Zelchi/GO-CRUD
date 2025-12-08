package route

import (
	"API_GO/internal/controller"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controller.Index)
}
