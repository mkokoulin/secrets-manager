// Package models includes models for the server
package models

import (
	"github.com/google/uuid"

	"encoding/json"
	"time"
)

// User storage structure
type User struct {
	ID        uuid.UUID `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

func (u User) MarshalJSON() ([]byte, error) {
	aliasValue := struct {
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
	}{
		ID:        u.ID.String(),
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
	}
	return json.Marshal(aliasValue)
}
