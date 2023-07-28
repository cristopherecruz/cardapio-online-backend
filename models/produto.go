package models

type Produto struct {
	Id           int           `json:"id"`
	Nome         string        `json:"nome"`
	Ingredientes []Ingrediente `json:"ingredientes"`
	Preco        float64       `json:"preco"`
}
