package database

import (
	"testing"

	"github.com/Luis-Andrei/api-users/models"
	"github.com/google/uuid"
)

func TestDatabase(t *testing.T) {
	db := NewDatabase()

	// Test Insert
	user := &models.User{
		ID:        uuid.New().String(),
		FirstName: "Test",
		LastName:  "User",
		Biography: "This is a test user biography with more than 20 characters",
	}
	insertedUser, err := db.Insert(user)
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}
	if insertedUser.FirstName != user.FirstName || insertedUser.LastName != user.LastName || insertedUser.Biography != user.Biography {
		t.Errorf("Inserted user does not match: got %v want %v", insertedUser, user)
	}

	// Test FindByID
	foundUser, err := db.FindByID(user.ID)
	if err != nil {
		t.Fatalf("FindByID failed: %v", err)
	}
	if foundUser.FirstName != user.FirstName || foundUser.LastName != user.LastName || foundUser.Biography != user.Biography {
		t.Errorf("Found user does not match: got %v want %v", foundUser, user)
	}

	// Test FindAll
	users := db.FindAll()
	if len(users) != 1 {
		t.Errorf("FindAll returned wrong number of users: got %v want %v", len(users), 1)
	}
	if users[0].FirstName != user.FirstName || users[0].LastName != user.LastName || users[0].Biography != user.Biography {
		t.Errorf("FindAll returned wrong user: got %v want %v", users[0], user)
	}

	// Test Update
	updatedUser := &models.User{
		FirstName: "Updated",
		LastName:  "User",
		Biography: "This is an updated test user biography with more than 20 characters",
	}
	updatedUser, err = db.Update(user.ID, updatedUser)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if updatedUser.FirstName != "Updated" || updatedUser.LastName != "User" || updatedUser.Biography != "This is an updated test user biography with more than 20 characters" {
		t.Errorf("Updated user does not match: got %v want %v", updatedUser, &models.User{
			FirstName: "Updated",
			LastName:  "User",
			Biography: "This is an updated test user biography with more than 20 characters",
		})
	}

	// Test Delete
	deletedUser, err := db.Delete(user.ID)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	if deletedUser.FirstName != "Updated" || deletedUser.LastName != "User" || deletedUser.Biography != "This is an updated test user biography with more than 20 characters" {
		t.Errorf("Deleted user does not match: got %v want %v", deletedUser, &models.User{
			FirstName: "Updated",
			LastName:  "User",
			Biography: "This is an updated test user biography with more than 20 characters",
		})
	}

	// Test FindByID after delete
	_, err = db.FindByID(user.ID)
	if err != ErrUserNotFound {
		t.Errorf("FindByID after delete returned wrong error: got %v want %v", err, ErrUserNotFound)
	}
}
