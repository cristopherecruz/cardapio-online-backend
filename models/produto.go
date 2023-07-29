package models

import "gorm.io/gorm"

type Produto struct {
	gorm.Model
	Nome         string        `json:"nome"`
	Ingredientes []Ingrediente `json:"ingredientes"`
	Preco        float64       `json:"preco"`
	ComboID      uint
}
