package main

import (
	"github.com/cristopherecruz/cardapio-online-backend/database"
	"github.com/cristopherecruz/cardapio-online-backend/routes"
)

func main() {

	database.ConectarBancoDados()
	routes.HandleRequest()

}
