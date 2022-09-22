package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	
	customerrors "github.com/mkokoulin/secrets-manager.git/internal/errors"
	"github.com/mkokoulin/secrets-manager.git/internal/helpers/encryptor"
	"github.com/mkokoulin/secrets-manager.git/internal/helpers/lunh"
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
	IsDeleted bool `json:"is_deleted"`
	Value map[string]string `json:"value"`
}

type RawSecretData struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID string `json:"user_id"`
	Type string `json:"type"`
	IsDeleted bool `json:"is_deleted"`
	Data   []byte `json:"secret_data"`
}

func (rsd *RawSecretData) MarshalJSON() ([]byte, error) {
	aliasValue := struct {
		CreatedAt string `json:"created_at"`
	}{
		CreatedAt:   rsd.CreatedAt.Format(time.RFC3339),
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
		UserID: secret.UserID,
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

func (rsd *RawSecretData) Encrypt() error {
	encryptValue, err := encryptor.Encrypt(key, nonce, rsd.Data)
	if err != nil {
		return err
	}

	rsd.Data = encryptValue

	return nil
}

func (rsd *RawSecretData) Decrypt() error {
	decryptValue, err := encryptor.Decrypt(key, nonce, rsd.Data)
	if err != nil {
		return err
	}

	rsd.Data = decryptValue

	return nil
}

func (s *SecretData) Validate() error {
	switch s.Type {
	case "binary":
		return s.validateBinary()
	case "login_password":
		return s.validateLoginPassword()
	case "credit_card":
		return s.validateCreditCard()
	case "string":
		return s.validateString()
	default:
		return customerrors.NewCustomError(errors.New("wrong type of secret"), "wrong type")
	}
}

func (sd *SecretData) validateBinary() error {
	return sd.checkUsefulData([]string{"binary"})
}

func (sd *SecretData) validateLoginPassword() error {
	return sd.checkUsefulData([]string{"login", "password"})
}

func (sd *SecretData) validateCreditCard() error {
	err := sd.checkUsefulData([]string{"card_number", "expired_date",
		"owner", "CVV"})
	if err != nil {
		return err
	}
	err = sd.checkFiledIsNumber([]string{"card_number", "CVV"})
	if err != nil {
		return err
	}
	cardNumber, _ := strconv.Atoi(sd.Value["card_number"])
	if !lunh.ValidLuhnNumber(cardNumber) {
		return customerrors.NewCustomError(
			errors.New("wrong credit card number"),
			"wrong credit card number")
	}
	return nil
}

func (sd *SecretData) validateString() error {
	return sd.checkUsefulData([]string{"string"})
}

func (sd *SecretData) checkFiledIsNumber(fields []string) error {
	var errorFields []string
	for _, field := range fields {
		if !isInt(field) {
			errorFields = append(errorFields, field)
		}
	}
	if len(errorFields) != 0 {
		text := fmt.Sprintf("this field must consist of numbers %v",
			strings.Trim(fmt.Sprint(errorFields), "[]"))
		return customerrors.NewCustomError(errors.New(text), text)
	}
	return nil
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func (sd *SecretData) checkUsefulData(fields []string) error {
	var missingFields []string
	for _, field := range fields {
		if _, ok := sd.Value[field]; !ok {
			missingFields = append(missingFields, field)
		}
	}
	if len(missingFields) != 0 {
		text := fmt.Sprintf("wrong format of data, missing fields %v",
			strings.Trim(fmt.Sprint(missingFields), "[]"))
		return customerrors.NewCustomError(errors.New(text), text)
	}
	return nil
}