package models

import "github.com/google/uuid"

type SecretType string

type SecretData struct {
	ID uuid.UUID
	Type SecretType
	Value []byte
}

type Secret struct {
	UserID uuid.UUID
	Data SecretData
}