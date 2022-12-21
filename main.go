package main

import (
	"net/http"
	routes "store/routes"
)

func main() {
	// Escuta o endereço inicial das rotas. No caso, o endereço '/' executando a função 'index'
	routes.LoadRoutes()

	// Provê um servidor http na porta 8000 com nenhuma configuração padrão
	http.ListenAndServe(":8000", nil)
}
