package route

import (
	"API_GO/internal/web/controller"
	"net/http"
)

func Load() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/new", controller.NewItem)
	http.HandleFunc("/insert", controller.Insert)
}
