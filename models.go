package main

import (
	"time"

	"github.com/firdauskosnin/RSSAgregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json: "id"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	Name      string    `json: "name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbuser database.User) User {
	return User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		APIKey:    dbuser.ApiKey,
	}
}
