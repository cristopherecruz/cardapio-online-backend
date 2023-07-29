package models

import "gorm.io/gorm"

type Combo struct {
	gorm.Model
	Nome      string    `json:"nome"`
	Descricao string    `json:"descricao"`
	Preco     float64   `json:"preco"`
	Produtos  []Produto `json:"produtos"`
}
