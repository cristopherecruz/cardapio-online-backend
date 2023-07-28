package models

type Ingrediente struct {
	Id        int    `json:"id"`
	Nome      string `json:"name"`
	ProdutoID string `json:"produto_id"`
}
