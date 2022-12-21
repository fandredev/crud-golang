package routes

import (
	"net/http"
	productController "store/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", productController.Index)
	http.HandleFunc("/new", productController.CreateProduct)
	http.HandleFunc("/insert", productController.InsertProduct)
	http.HandleFunc("/delete", productController.DeleteProduct)
	http.HandleFunc("/edit", productController.EditProduct)
	http.HandleFunc("/update", productController.UpdateProduct)
}
