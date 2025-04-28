package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/yourusername/crud-api/models"
)

var (
	itens = make(map[string]models.Item)
	mutex = &sync.RWMutex{}
)

// CriarItem lida com requisições POST para criar um novo item
func CriarItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Erro ao decodificar o JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	itens[item.ID] = item
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// ObterItem lida com requisições GET para recuperar um item por ID
func ObterItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}

	mutex.RLock()
	item, existe := itens[id]
	mutex.RUnlock()

	if !existe {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// AtualizarItem lida com requisições PUT para atualizar um item existente
func AtualizarItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Erro ao decodificar o JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	if _, existe := itens[id]; !existe {
		mutex.Unlock()
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	itens[id] = item
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// DeletarItem lida com requisições DELETE para remover um item
func DeletarItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	if _, existe := itens[id]; !existe {
		mutex.Unlock()
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	delete(itens, id)
	mutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

// ObterTodosItens lida com requisições GET para recuperar todos os itens
func ObterTodosItens(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	todosItens := make([]models.Item, 0, len(itens))
	for _, item := range itens {
		todosItens = append(todosItens, item)
	}
	mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todosItens)
}
