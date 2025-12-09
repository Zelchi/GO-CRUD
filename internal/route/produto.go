package route

import (
	"API_GO/internal/web/controller"
	"net/http"
)

func LoadFront() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/new", controller.NewItem)
}

func LoadBack() {
	http.HandleFunc("/api/insert", controller.Insert)
}
