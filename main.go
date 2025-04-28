package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/crud-api/handlers"
)

func main() {
	router := mux.NewRouter()

	// Define as rotas da API
	router.HandleFunc("/itens", handlers.CriarItem).Methods("POST")
	router.HandleFunc("/itens", handlers.ObterTodosItens).Methods("GET")
	router.HandleFunc("/itens/{id}", handlers.ObterItem).Methods("GET")
	router.HandleFunc("/itens/{id}", handlers.AtualizarItem).Methods("PUT")
	router.HandleFunc("/itens/{id}", handlers.DeletarItem).Methods("DELETE")

	// Inicia o servidor
	log.Println("Servidor iniciando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
