package database

import (
	"os"
	"testing"

	"github.com/Luis-Andrei/api-users/models"
)

func setupTestDB(t *testing.T) *PostgresDB {
	host := os.Getenv("TEST_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("TEST_DB_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("TEST_DB_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("TEST_DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	dbname := os.Getenv("TEST_DB_NAME")
	if dbname == "" {
		dbname = "bank_test"
	}

	db, err := NewPostgresDB(host, port, user, password, dbname)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	if err := db.InitTables(); err != nil {
		t.Fatalf("Failed to initialize tables: %v", err)
	}

	// Limpa a tabela de clientes antes de cada teste
	postgresDB, ok := db.(*PostgresDB)
	if !ok {
		t.Fatal("Failed to convert database to PostgresDB type")
	}
	_, err = postgresDB.db.Exec("DELETE FROM clients")
	if err != nil {
		t.Fatalf("Failed to clean clients table: %v", err)
	}

	return postgresDB
}

func TestPostgresDB_PersonalClient(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Cria um cliente pessoa física
	client := models.NewPersonalClient("John Doe", "123.456.789-00", 2000.0)

	// Testa a criação do cliente
	err := db.CreatePersonalClient(client)
	if err != nil {
		t.Fatalf("Failed to create personal client: %v", err)
	}

	// Testa a busca do cliente
	foundClient, err := db.GetClient(client.ID)
	if err != nil {
		t.Fatalf("Failed to get client: %v", err)
	}

	personalClient, ok := foundClient.(*models.PersonalClient)
	if !ok {
		t.Fatal("Retrieved client is not a personal client")
	}

	if personalClient.Name != client.Name || personalClient.CPF != client.CPF || personalClient.Balance != client.Balance {
		t.Errorf("Retrieved client does not match created client")
	}

	// Testa o saque
	err = personalClient.Withdraw(500.0)
	if err != nil {
		t.Fatalf("Failed to withdraw: %v", err)
	}

	// Testa a atualização do cliente
	err = db.UpdateClient(personalClient)
	if err != nil {
		t.Fatalf("Failed to update client: %v", err)
	}

	// Verifica se o saldo foi atualizado
	updatedClient, err := db.GetClient(client.ID)
	if err != nil {
		t.Fatalf("Failed to get updated client: %v", err)
	}

	if updatedClient.GetBalance() != 1500.0 {
		t.Errorf("Expected balance of 1500.0, got %v", updatedClient.GetBalance())
	}
}

func TestPostgresDB_CorporateClient(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Cria um cliente pessoa jurídica
	client := models.NewCorporateClient("ACME Corp", "12.345.678/0001-00", 10000.0)

	// Testa a criação do cliente
	err := db.CreateCorporateClient(client)
	if err != nil {
		t.Fatalf("Failed to create corporate client: %v", err)
	}

	// Testa a busca do cliente
	foundClient, err := db.GetClient(client.ID)
	if err != nil {
		t.Fatalf("Failed to get client: %v", err)
	}

	corporateClient, ok := foundClient.(*models.CorporateClient)
	if !ok {
		t.Fatal("Retrieved client is not a corporate client")
	}

	if corporateClient.Name != client.Name || corporateClient.CNPJ != client.CNPJ || corporateClient.Balance != client.Balance {
		t.Errorf("Retrieved client does not match created client")
	}

	// Testa o saque
	err = corporateClient.Withdraw(3000.0)
	if err != nil {
		t.Fatalf("Failed to withdraw: %v", err)
	}

	// Testa a atualização do cliente
	err = db.UpdateClient(corporateClient)
	if err != nil {
		t.Fatalf("Failed to update client: %v", err)
	}

	// Verifica se o saldo foi atualizado
	updatedClient, err := db.GetClient(client.ID)
	if err != nil {
		t.Fatalf("Failed to get updated client: %v", err)
	}

	if updatedClient.GetBalance() != 7000.0 {
		t.Errorf("Expected balance of 7000.0, got %v", updatedClient.GetBalance())
	}
}

func TestPostgresDB_ListClients(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Cria um cliente pessoa física
	personalClient := models.NewPersonalClient("John Doe", "123.456.789-00", 2000.0)
	err := db.CreatePersonalClient(personalClient)
	if err != nil {
		t.Fatalf("Failed to create personal client: %v", err)
	}

	// Cria um cliente pessoa jurídica
	corporateClient := models.NewCorporateClient("ACME Corp", "12.345.678/0001-00", 10000.0)
	err = db.CreateCorporateClient(corporateClient)
	if err != nil {
		t.Fatalf("Failed to create corporate client: %v", err)
	}

	// Lista todos os clientes
	clients, err := db.ListClients()
	if err != nil {
		t.Fatalf("Failed to list clients: %v", err)
	}

	if len(clients) != 2 {
		t.Errorf("Expected 2 clients, got %d", len(clients))
	}

	// Verifica se os clientes estão corretos
	var foundPersonal, foundCorporate bool
	for _, client := range clients {
		switch c := client.(type) {
		case *models.PersonalClient:
			if c.Name == personalClient.Name && c.CPF == personalClient.CPF {
				foundPersonal = true
			}
		case *models.CorporateClient:
			if c.Name == corporateClient.Name && c.CNPJ == corporateClient.CNPJ {
				foundCorporate = true
			}
		}
	}

	if !foundPersonal {
		t.Error("Personal client not found in list")
	}
	if !foundCorporate {
		t.Error("Corporate client not found in list")
	}
}
