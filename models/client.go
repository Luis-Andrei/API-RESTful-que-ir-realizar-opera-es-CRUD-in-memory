package models

import (
	"errors"
	"time"
)

// Errors
var (
	ErrInsufficientFunds = errors.New("saldo insuficiente")
	ErrInvalidAmount     = errors.New("valor inválido")
	ErrWithdrawLimit     = errors.New("limite de saque excedido")
)

// Transaction representa uma transação bancária
type Transaction struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"` // "withdrawal" ou "deposit"
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// Client é a interface que define os métodos que um cliente deve implementar
type Client interface {
	// Withdraw realiza um saque na conta do cliente
	Withdraw(amount float64) error

	// GetStatement retorna o extrato das transações do cliente
	GetStatement() []Transaction

	// GetBalance retorna o saldo atual do cliente
	GetBalance() float64

	// GetWithdrawLimit retorna o limite de saque do cliente
	GetWithdrawLimit() float64
}

// BaseClient contém os campos comuns entre pessoa física e jurídica
type BaseClient struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

// PersonalClient representa uma pessoa física
type PersonalClient struct {
	BaseClient
	CPF string `json:"cpf"`
}

// CorporateClient representa uma pessoa jurídica
type CorporateClient struct {
	BaseClient
	CNPJ string `json:"cnpj"`
}

// Constantes para limites de saque
const (
	PersonalClientWithdrawLimit  = 1000.0
	CorporateClientWithdrawLimit = 5000.0
)
