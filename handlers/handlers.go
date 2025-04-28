package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Luis-Andrei/api-users/database"
	"github.com/Luis-Andrei/api-users/models"
	"github.com/gorilla/mux"
)

type Handler struct {
	db database.Database
}

func NewHandler(db database.Database) *Handler {
	return &Handler{db: db}
}

type CreatePersonalClientRequest struct {
	Name           string  `json:"name"`
	CPF            string  `json:"cpf"`
	InitialBalance float64 `json:"initial_balance"`
}

type CreateCorporateClientRequest struct {
	Name           string  `json:"name"`
	CNPJ           string  `json:"cnpj"`
	InitialBalance float64 `json:"initial_balance"`
}

type WithdrawRequest struct {
	Amount float64 `json:"amount"`
}

func (h *Handler) CreatePersonalClient(w http.ResponseWriter, r *http.Request) {
	var req CreatePersonalClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client := models.NewPersonalClient(req.Name, req.CPF, req.InitialBalance)
	if err := h.db.CreatePersonalClient(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func (h *Handler) CreateCorporateClient(w http.ResponseWriter, r *http.Request) {
	var req CreateCorporateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client := models.NewCorporateClient(req.Name, req.CNPJ, req.InitialBalance)
	if err := h.db.CreateCorporateClient(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func (h *Handler) GetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	client, err := h.db.GetClient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func (h *Handler) ListClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.db.ListClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req WithdrawRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client, err := h.db.GetClient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := client.Withdraw(req.Amount); err != nil {
		switch err {
		case models.ErrInvalidAmount:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case models.ErrInsufficientFunds:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case models.ErrWithdrawLimit:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := h.db.UpdateClient(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func (h *Handler) GetStatement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	client, err := h.db.GetClient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client.GetStatement())
}
