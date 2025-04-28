package models

// Item representa um produto na nossa API
type Item struct {
	ID    string  `json:"id"`    // Identificador único do item
	Name  string  `json:"name"`  // Nome do item
	Price float64 `json:"price"` // Preço do item
}
