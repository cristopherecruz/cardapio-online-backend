package database

import (
	"fmt"
	"github.com/cristopherecruz/cardapio-online-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB
var Err error

func ConectarBancoDados() {

	connStr := "host=localhost user=postgres password=root dbname=cardapio port=5432 sslmode=disable TimeZone=America/SAO_PAULO"

	Db, Err = gorm.Open(postgres.Open(connStr))
	if Err != nil {
		log.Panic(Err.Error())
	}

	Db.AutoMigrate(&models.Produto{})
	Db.AutoMigrate(&models.Ingrediente{})
	Db.AutoMigrate(&models.Combo{})

	//PopularDados(Db)

}

func PopularDados(db *gorm.DB) {
	// Crie um combo com produtos e ingredientes associados
	ingrediente1 := models.Ingrediente{Nome: "Ingrediente 1"}
	ingrediente2 := models.Ingrediente{Nome: "Ingrediente 2"}

	produto1 := models.Produto{Nome: "Produto 1", Preco: 10.50}
	produto1.Ingredientes = append(produto1.Ingredientes, ingrediente1, ingrediente2)

	produto2 := models.Produto{Nome: "Produto 2", Preco: 15.75}
	produto2.Ingredientes = append(produto2.Ingredientes, ingrediente2)

	combo := models.Combo{Nome: "Combo 1", Descricao: "Combo com produtos e ingredientes", Preco: 25.99}
	combo.Produtos = append(combo.Produtos, produto1, produto2)

	// Salvar o combo com produtos e ingredientes no banco de dados
	db.Create(&combo)

	fmt.Println("Combo criado com sucesso!")

	// Recuperar o combo com produtos e ingredientes do banco de dados
	var retrievedCombo models.Combo
	db.Preload("Produtos.Ingredientes").First(&retrievedCombo, combo.ID)

	fmt.Println("Combo:", retrievedCombo.Nome)
	fmt.Println("Descrição:", retrievedCombo.Descricao)
	fmt.Println("Preço:", retrievedCombo.Preco)

	fmt.Println("Produtos:")
	for _, produto := range retrievedCombo.Produtos {
		fmt.Printf("  Produto: %s, Preço: %.2f\n", produto.Nome, produto.Preco)

		fmt.Println("  Ingredientes:")
		for _, ingrediente := range produto.Ingredientes {
			fmt.Printf("    Ingrediente: %s\n", ingrediente.Nome)
		}
	}

}
