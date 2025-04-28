package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Luis-Andrei/api-users/database"
	"github.com/Luis-Andrei/api-users/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Configuração do banco de dados
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "bank"
	}

	// Conecta ao banco de dados
	db, err := database.NewPostgresDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Inicializa as tabelas
	if err := db.InitTables(); err != nil {
		log.Fatalf("Erro ao inicializar tabelas: %v", err)
	}

	// Cria uma nova instância do handler
	handler := handlers.NewHandler(db)

	// Cria um novo router
	router := mux.NewRouter()

	// Define as rotas da API
	router.HandleFunc("/api/clients/personal", handler.CreatePersonalClient).Methods("POST")
	router.HandleFunc("/api/clients/corporate", handler.CreateCorporateClient).Methods("POST")
	router.HandleFunc("/api/clients", handler.ListClients).Methods("GET")
	router.HandleFunc("/api/clients/{id}", handler.GetClient).Methods("GET")
	router.HandleFunc("/api/clients/{id}/withdraw", handler.Withdraw).Methods("POST")
	router.HandleFunc("/api/clients/{id}/statement", handler.GetStatement).Methods("GET")

	// Inicia o servidor
	log.Println("Servidor iniciando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
