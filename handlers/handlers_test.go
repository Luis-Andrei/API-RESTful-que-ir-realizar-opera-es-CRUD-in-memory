package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Luis-Andrei/api-users/database"
	"github.com/Luis-Andrei/api-users/models"
	"github.com/gorilla/mux"
)

func setupTestHandler(t *testing.T) *Handler {
	db := &database.MockDB{
		OnCreatePersonalClient: func(client *models.PersonalClient) error {
			return nil
		},
		OnCreateCorporateClient: func(client *models.CorporateClient) error {
			return nil
		},
		OnGetClient: func(id string) (models.Client, error) {
			if id == "123" {
				return models.NewPersonalClient("John Doe", "123.456.789-00", 2000.0), nil
			}
			return nil, nil
		},
		OnUpdateClient: func(client models.Client) error {
			return nil
		},
		OnListClients: func() ([]models.Client, error) {
			return nil, nil
		},
	}

	return NewHandler(db)
}

func TestCreatePersonalClient(t *testing.T) {
	handler := setupTestHandler(t)

	reqBody := CreatePersonalClientRequest{
		Name:           "John Doe",
		CPF:            "123.456.789-00",
		InitialBalance: 2000.0,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/api/clients/personal", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.CreatePersonalClient(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestCreateCorporateClient(t *testing.T) {
	handler := setupTestHandler(t)

	reqBody := CreateCorporateClientRequest{
		Name:           "ACME Corp",
		CNPJ:           "12.345.678/0001-00",
		InitialBalance: 10000.0,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/api/clients/corporate", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.CreateCorporateClient(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetClient(t *testing.T) {
	handler := setupTestHandler(t)

	req := httptest.NewRequest("GET", "/api/clients/123", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "123"})
	w := httptest.NewRecorder()

	handler.GetClient(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestListClients(t *testing.T) {
	handler := setupTestHandler(t)

	req := httptest.NewRequest("GET", "/api/clients", nil)
	w := httptest.NewRecorder()

	handler.ListClients(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestWithdraw(t *testing.T) {
	handler := setupTestHandler(t)

	reqBody := WithdrawRequest{
		Amount: 500.0,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/api/clients/123/withdraw", bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": "123"})
	w := httptest.NewRecorder()

	handler.Withdraw(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetStatement(t *testing.T) {
	handler := setupTestHandler(t)

	req := httptest.NewRequest("GET", "/api/clients/123/statement", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "123"})
	w := httptest.NewRecorder()

	handler.GetStatement(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
