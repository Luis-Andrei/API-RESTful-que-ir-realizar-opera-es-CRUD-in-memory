package models

import (
	"time"

	"github.com/google/uuid"
)

// NewPersonalClient cria um novo cliente pessoa física
func NewPersonalClient(name, cpf string, initialBalance float64) *PersonalClient {
	return &PersonalClient{
		BaseClient: BaseClient{
			ID:           uuid.New().String(),
			Name:         name,
			Balance:      initialBalance,
			Transactions: make([]Transaction, 0),
		},
		CPF: cpf,
	}
}

// NewCorporateClient cria um novo cliente pessoa jurídica
func NewCorporateClient(name, cnpj string, initialBalance float64) *CorporateClient {
	return &CorporateClient{
		BaseClient: BaseClient{
			ID:           uuid.New().String(),
			Name:         name,
			Balance:      initialBalance,
			Transactions: make([]Transaction, 0),
		},
		CNPJ: cnpj,
	}
}

// Implementação dos métodos para PersonalClient

func (c *PersonalClient) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > c.GetWithdrawLimit() {
		return ErrWithdrawLimit
	}
	if amount > c.Balance {
		return ErrInsufficientFunds
	}

	c.Balance -= amount
	c.Transactions = append(c.Transactions, Transaction{
		ID:          uuid.New().String(),
		Amount:      amount,
		Type:        "withdrawal",
		Description: "Saque em dinheiro",
		CreatedAt:   time.Now(),
	})

	return nil
}

func (c *PersonalClient) GetStatement() []Transaction {
	return c.Transactions
}

func (c *PersonalClient) GetBalance() float64 {
	return c.Balance
}

func (c *PersonalClient) GetWithdrawLimit() float64 {
	return PersonalClientWithdrawLimit
}

// Implementação dos métodos para CorporateClient

func (c *CorporateClient) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > c.GetWithdrawLimit() {
		return ErrWithdrawLimit
	}
	if amount > c.Balance {
		return ErrInsufficientFunds
	}

	c.Balance -= amount
	c.Transactions = append(c.Transactions, Transaction{
		ID:          uuid.New().String(),
		Amount:      amount,
		Type:        "withdrawal",
		Description: "Saque em dinheiro",
		CreatedAt:   time.Now(),
	})

	return nil
}

func (c *CorporateClient) GetStatement() []Transaction {
	return c.Transactions
}

func (c *CorporateClient) GetBalance() float64 {
	return c.Balance
}

func (c *CorporateClient) GetWithdrawLimit() float64 {
	return CorporateClientWithdrawLimit
}
