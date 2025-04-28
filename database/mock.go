package database

import "github.com/Luis-Andrei/api-users/models"

// MockDB é uma implementação mock do banco de dados para testes
type MockDB struct {
	OnCreatePersonalClient  func(client *models.PersonalClient) error
	OnCreateCorporateClient func(client *models.CorporateClient) error
	OnGetClient             func(id string) (models.Client, error)
	OnUpdateClient          func(client models.Client) error
	OnListClients           func() ([]models.Client, error)
}

func (m *MockDB) CreatePersonalClient(client *models.PersonalClient) error {
	if m.OnCreatePersonalClient != nil {
		return m.OnCreatePersonalClient(client)
	}
	return nil
}

func (m *MockDB) CreateCorporateClient(client *models.CorporateClient) error {
	if m.OnCreateCorporateClient != nil {
		return m.OnCreateCorporateClient(client)
	}
	return nil
}

func (m *MockDB) GetClient(id string) (models.Client, error) {
	if m.OnGetClient != nil {
		return m.OnGetClient(id)
	}
	return nil, nil
}

func (m *MockDB) UpdateClient(client models.Client) error {
	if m.OnUpdateClient != nil {
		return m.OnUpdateClient(client)
	}
	return nil
}

func (m *MockDB) ListClients() ([]models.Client, error) {
	if m.OnListClients != nil {
		return m.OnListClients()
	}
	return nil, nil
}

func (m *MockDB) Close() error {
	return nil
}

func (m *MockDB) InitTables() error {
	return nil
}
