package models

import (
	"errors"
	"strings"
)

// User representa a estrutura de um usuário
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

// Validate verifica se os campos do usuário são válidos
func (u *User) Validate() error {
	if len(u.FirstName) < 2 || len(u.FirstName) > 20 {
		return errors.New("first_name deve ter entre 2 e 20 caracteres")
	}
	if len(u.LastName) < 2 || len(u.LastName) > 20 {
		return errors.New("last_name deve ter entre 2 e 20 caracteres")
	}
	if len(u.Biography) < 20 || len(u.Biography) > 450 {
		return errors.New("biography deve ter entre 20 e 450 caracteres")
	}
	return nil
}

// NewUser cria um novo usuário com os dados fornecidos
func NewUser(firstName, lastName, biography string) (*User, error) {
	user := &User{
		FirstName: strings.TrimSpace(firstName),
		LastName:  strings.TrimSpace(lastName),
		Biography: strings.TrimSpace(biography),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}
