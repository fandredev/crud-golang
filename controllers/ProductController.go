package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	products "store/models"
	"strconv"
)

// Carrega todos os templates da pasta 'templates/' com extensão final *html
var loadTemplates = template.Must(template.ParseGlob("templates/*.html"))

const statusMovedPermanently = http.StatusMovedPermanently

func Index(w http.ResponseWriter, r *http.Request) {
	products := products.FindAllProductsFromDatabase()
	// Carrega o template Index com os produtos
	loadTemplates.ExecuteTemplate(w, "Index", products)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	loadTemplates.ExecuteTemplate(w, "CreateNewProduct", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nameProduct := r.FormValue("name")
		description := r.FormValue("description")
		price, errConvertedPrice := strconv.ParseFloat(r.FormValue("price"), 64)
		quantity, errConvertedQuantity := strconv.Atoi(r.FormValue("quantity"))

		if errConvertedPrice != nil {
			fmt.Println("Erro na conversão do preço para ponto flutuante.")
		}

		if errConvertedQuantity != nil {
			fmt.Println("Erro na conversão da quantidade para númerico.")
		}

		products.CreateNewProductFromDatabase(nameProduct, description, price, quantity)
		http.Redirect(w, r, "/", statusMovedPermanently)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idProductFromURL := r.URL.Query().Get("id")

	products.DeleteProductFromDatabase(idProductFromURL)
	http.Redirect(w, r, "/", statusMovedPermanently)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	idProductFromURL := r.URL.Query().Get("id")
	productFiltered := products.FindProductById(idProductFromURL)

	loadTemplates.ExecuteTemplate(w, "EditProduct", productFiltered)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, errIdConvertedToInt := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		description := r.FormValue("description")
		quantity, errConvertedQuantityToInt := strconv.Atoi(r.FormValue("quantity"))
		price, errConvertedPriceToFloat := strconv.ParseFloat(r.FormValue("price"), 64)

		if errIdConvertedToInt != nil {
			log.Panicln("Erro na conversão do id do produto para inteiro.")
		}

		if errConvertedQuantityToInt != nil {
			log.Panicln("Erro na conversão de quantidade de produtos para inteiro.")
		}

		if errConvertedPriceToFloat != nil {
			log.Panicln("Erro na conversão de preço do produto para ponto flutuante.")
		}

		products.UpdateProduct(id, quantity, name, description, price)
	}
	http.Redirect(w, r, "/", statusMovedPermanently)
}
