package database

import (
	"errors"
	"sync"

	"github.com/Luis-Andrei/api-users/models"
	"github.com/google/uuid"
)

// Database é a interface que define os métodos que um banco de dados deve implementar
type Database interface {
	CreatePersonalClient(client *models.PersonalClient) error
	CreateCorporateClient(client *models.CorporateClient) error
	GetClient(id string) (models.Client, error)
	UpdateClient(client models.Client) error
	ListClients() ([]models.Client, error)
	Close() error
	InitTables() error
}

// Database representa o banco de dados em memória
type DatabaseStruct struct {
	users map[string]*models.User
	mutex sync.RWMutex
}

// NewDatabase cria uma nova instância do banco de dados
func NewDatabase() *DatabaseStruct {
	return &DatabaseStruct{
		users: make(map[string]*models.User),
	}
}

// FindAll retorna todos os usuários
func (db *DatabaseStruct) FindAll() []*models.User {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	users := make([]*models.User, 0, len(db.users))
	for _, user := range db.users {
		users = append(users, user)
	}
	return users
}

// FindByID retorna um usuário pelo ID
func (db *DatabaseStruct) FindByID(id string) (*models.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	user, exists := db.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// Insert insere um novo usuário
func (db *DatabaseStruct) Insert(user *models.User) (*models.User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	user.ID = uuid.New().String()
	db.users[user.ID] = user
	return user, nil
}

// Update atualiza um usuário existente
func (db *DatabaseStruct) Update(id string, user *models.User) (*models.User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.users[id]; !exists {
		return nil, ErrUserNotFound
	}

	user.ID = id
	db.users[id] = user
	return user, nil
}

// Delete remove um usuário
func (db *DatabaseStruct) Delete(id string) (*models.User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	user, exists := db.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	delete(db.users, id)
	return user, nil
}

// Erros do banco de dados
var (
	ErrUserNotFound = errors.New("usuário não encontrado")
)
