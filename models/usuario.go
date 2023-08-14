package models

import (
	"github.com/cristopherecruz/cardapio-online-backend/config/seguranca"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type Usuario struct {
	gorm.Model
	NomeUsuario string `json:"usuario"`
	Senha       string `json:"senha"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ChecarLogin(username string, password string, dados *gorm.DB) (string, error) {

	var err error

	u := Usuario{}

	err = dados.Model(Usuario{}).Where("nome_usuario = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Senha)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := seguranca.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *Usuario) BeforeSave(*gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Senha = string(hashedPassword)
	u.NomeUsuario = html.EscapeString(strings.TrimSpace(u.NomeUsuario))

	return nil
}
