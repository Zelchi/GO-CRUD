package route

import (
	"API_GO/internal/web/controller"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.HomePage)
	http.HandleFunc("/produto/new", controller.AddItemPage)
	http.HandleFunc("/produto/edit", controller.EditPage)

	http.HandleFunc("/produto/insert", controller.Insert)
	http.HandleFunc("/produto/update", controller.Update)
	http.HandleFunc("/produto/delete", controller.Delete)
}
