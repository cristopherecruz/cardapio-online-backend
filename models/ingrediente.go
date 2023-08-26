package models

import "gorm.io/gorm"

type Ingrediente struct {
	gorm.Model
	Nome      string `json:"nome"`
	Marca     string `json:"marca"`
	ProdutoID uint
}
