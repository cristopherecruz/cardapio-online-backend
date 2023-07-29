package models

import "gorm.io/gorm"

type Ingrediente struct {
	gorm.Model
	Nome      string `json:"nome"`
	ProdutoID uint
}
