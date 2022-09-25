// Package encryptor required for encryption and decryption
package encryptor

import (
	"crypto/aes"
	"crypto/cipher"

	customerrors "github.com/mkokoulin/secrets-manager.git/server/internal/errors"

)

// Encrypt function for encrypting information
func Encrypt(key, nonce,  data []byte) ([]byte, error) {
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return nil, customerrors.NewCustomError(err, "error with encrypt")
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, customerrors.NewCustomError(err, "error with encrypt")
	}

	encryptData := aesgcm.Seal(nil, nonce, data, nil)
		
	return encryptData, nil
}

// Decrypt function for decrypting information
func Decrypt(key, nonce,  data []byte) ([]byte, error) {
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	decryptData, err := aesgcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}

	return decryptData, nil
}	