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

type Feed struct {
	ID        uuid.UUID `json: "id"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	Name      string    `json: "name"`
	Url       string    `json: "url"`
	UserID    uuid.UUID `json: "user_id"`
}

func databseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.ID,
	}
}
