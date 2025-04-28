package models

import (
	"testing"
)

func TestPersonalClient(t *testing.T) {
	client := NewPersonalClient("John Doe", "123.456.789-00", 2000.0)

	// Test initial state
	if client.GetBalance() != 2000.0 {
		t.Errorf("Expected initial balance of 2000.0, got %v", client.GetBalance())
	}
	if client.GetWithdrawLimit() != PersonalClientWithdrawLimit {
		t.Errorf("Expected withdraw limit of %v, got %v", PersonalClientWithdrawLimit, client.GetWithdrawLimit())
	}

	// Test valid withdrawal
	err := client.Withdraw(500.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if client.GetBalance() != 1500.0 {
		t.Errorf("Expected balance of 1500.0, got %v", client.GetBalance())
	}

	// Test withdrawal above limit
	err = client.Withdraw(1500.0)
	if err != ErrWithdrawLimit {
		t.Errorf("Expected ErrWithdrawLimit, got %v", err)
	}

	// Test invalid amount
	err = client.Withdraw(-100.0)
	if err != ErrInvalidAmount {
		t.Errorf("Expected ErrInvalidAmount, got %v", err)
	}

	// Test insufficient funds
	err = client.Withdraw(2000.0)
	if err != ErrWithdrawLimit {
		t.Errorf("Expected ErrWithdrawLimit, got %v", err)
	}

	// Test statement
	transactions := client.GetStatement()
	if len(transactions) != 1 {
		t.Errorf("Expected 1 transaction, got %v", len(transactions))
	}
	if transactions[0].Amount != 500.0 {
		t.Errorf("Expected transaction amount of 500.0, got %v", transactions[0].Amount)
	}
	if transactions[0].Type != "withdrawal" {
		t.Errorf("Expected transaction type 'withdrawal', got %v", transactions[0].Type)
	}
}

func TestCorporateClient(t *testing.T) {
	client := NewCorporateClient("ACME Corp", "12.345.678/0001-00", 10000.0)

	// Test initial state
	if client.GetBalance() != 10000.0 {
		t.Errorf("Expected initial balance of 10000.0, got %v", client.GetBalance())
	}
	if client.GetWithdrawLimit() != CorporateClientWithdrawLimit {
		t.Errorf("Expected withdraw limit of %v, got %v", CorporateClientWithdrawLimit, client.GetWithdrawLimit())
	}

	// Test valid withdrawal
	err := client.Withdraw(3000.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if client.GetBalance() != 7000.0 {
		t.Errorf("Expected balance of 7000.0, got %v", client.GetBalance())
	}

	// Test withdrawal above limit
	err = client.Withdraw(6000.0)
	if err != ErrWithdrawLimit {
		t.Errorf("Expected ErrWithdrawLimit, got %v", err)
	}

	// Test invalid amount
	err = client.Withdraw(-100.0)
	if err != ErrInvalidAmount {
		t.Errorf("Expected ErrInvalidAmount, got %v", err)
	}

	// Test insufficient funds
	err = client.Withdraw(8000.0)
	if err != ErrWithdrawLimit {
		t.Errorf("Expected ErrWithdrawLimit, got %v", err)
	}

	// Test statement
	transactions := client.GetStatement()
	if len(transactions) != 1 {
		t.Errorf("Expected 1 transaction, got %v", len(transactions))
	}
	if transactions[0].Amount != 3000.0 {
		t.Errorf("Expected transaction amount of 3000.0, got %v", transactions[0].Amount)
	}
	if transactions[0].Type != "withdrawal" {
		t.Errorf("Expected transaction type 'withdrawal', got %v", transactions[0].Type)
	}
}
