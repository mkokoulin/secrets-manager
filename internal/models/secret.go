package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	customerrors "github.com/mkokoulin/secrets-manager.git/internal/errors"
	"github.com/mkokoulin/secrets-manager.git/internal/helpers/encryptor"
)

var (
	key   = []byte{240, 43, 127, 3, 22, 181, 93, 105, 162, 19, 180, 125, 207, 77, 209, 70}
	nonce = []byte{161, 154, 38, 17, 9, 137, 119, 105, 204, 99, 67, 14}
)

type Secret struct {
	UserID string `json:"user_id"`
	SecretID string 
	Data SecretData `db:"secrets_data" gorm:"foreignKey:SecretID"`
}

type SecretData struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Type string `json:"type"`
	Value map[string]string `json:"value"`
}

type RawSecretData struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID string `json:"user_id"`
	Type string `json:"type"`
	Data   []byte `json:"secret_data"`
}

func (s *RawSecretData) MarshalJSON() ([]byte, error) {
	aliasValue := struct {
		CreatedAt string `json:"created_at"`
	}{
		CreatedAt:   s.CreatedAt.Format(time.RFC3339),
	}
	return json.Marshal(aliasValue)
}

func NewRawSecretData(secret Secret) (*RawSecretData, error) {
	err := secret.Data.Validate()
	if err != nil {
		return nil, err
	}

	value, err := json.Marshal(secret.Data.Value)
	if err != nil {
		return nil, err
	}

	data := RawSecretData{
		ID: uuid.New().String(),
		CreatedAt: time.Now(),
		Type: secret.Data.Type,
		Data: value,
	}

	err = data.Encrypt()
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *RawSecretData) Encrypt() error {
	encryptValue, err := encryptor.Encrypt(key, nonce, s.Data)
	if err != nil {
		return err
	}

	s.Data = encryptValue

	return nil
}

func (s *RawSecretData) Decrypt() error {
	decryptValue, err := encryptor.Decrypt(key, nonce, s.Data)
	if err != nil {
		return err
	}

	s.Data = decryptValue

	return nil
}

// func (s *SecretData) MarshalJSON() ([]byte, error) {
// 	err := s.Validate()
// 	if err != nil {
// 		return nil, err
// 	}

// 	value, err := json.Marshal(s.Value)
// 	if err != nil {
// 		return nil, err
// 	}

// 	encryptValue, err := encryptor.Encrypt(key, nonce, value)
// 	if err != nil {
// 		return nil, err
// 	}

// 	aliasValue := struct {
// 		CreatedAt string `json:"created_at"`
// 		Value []byte `json:"value"`
// 	}{
// 		CreatedAt:   s.CreatedAt.Format(time.RFC3339),
// 		Value: encryptValue,
// 	}
	
// 	return json.Marshal(aliasValue)
// }

// func (s *SecretData) UnmarshalJSON(b []byte) error {
// 	sd := &struct {
// 		Value []byte `json:"value"`
// 	}{}

// 	err := json.Unmarshal(b, &sd)
// 	if err != nil {
// 		return err
// 	}

// 	decryptValue, err := encryptor.Decrypt(key, nonce, sd.Value)
// 	if err != nil {
// 		return err
// 	}

// 	var value map[string]string

// 	err = json.Unmarshal(decryptValue, &value)
// 	if err != nil {
// 		return err
// 	}

// 	// s.Value = value


// 	return nil
// }

func (s *SecretData) Validate() error {
	switch s.Type {
	case "binary":
		return nil
	case "login_password":
		return nil
	case "credit_card":
		return nil
	case "string":
		return nil
	default:
		return customerrors.NewCustomError(errors.New("wrong type of secret"), "wrong type")
	}
}