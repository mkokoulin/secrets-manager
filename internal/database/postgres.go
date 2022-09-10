package database

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	customerrors "github.com/mkokoulin/secrets-manager.git/internal/errors"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
)

type PostgresDatabase struct {
	conn *gorm.DB
}

func NewPostgresDatabase(conn *gorm.DB) *PostgresDatabase {
	return &PostgresDatabase{
		conn: conn,
	}
}

func (pd *PostgresDatabase) CreateUser(ctx context.Context, user models.User) error {
	err := pd.conn.Transaction(func(tx *gorm.DB) error {
		var exists bool

		err := tx.Model(&models.User{}).Select("count(*) > 0").Where("login = ?", user.Login).Find(&exists).Error
		if err != nil {
			return customerrors.NewCustomError(err, "an unknown error occurred during checking the user")
		}

		if exists {
			return customerrors.NewCustomError(errors.New(""), "user already exists")
		}


		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return customerrors.NewCustomError(err, "an unknown error occurred during generation password")
		}

		user.Password = string(hash)

		if err := tx.Create(&user).Error; err != nil {
			return customerrors.NewCustomError(err, "an unknown error occurred during user creation")
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (pd *PostgresDatabase) CheckUserPassword(ctx context.Context, user models.User) (string, error) {
	var result models.User

	err := pd.conn.Model(&models.User{}).Where("login = ?", user.Login).First(&result).Error
	if err != nil {
		return result.Login, customerrors.NewCustomError(err, "an unknown error occurred during getting a user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
        return result.Login, customerrors.NewCustomError(err, "an unknown error occurred during generation password")
    }

	return result.Login, nil
}

func (pd *PostgresDatabase) DeleteUser(ctx context.Context, userID string) error {	
	err := pd.conn.Model(&models.User{}).Where("id=?", userID).Update("is_deleted", true).Error
	if err != nil {
		return customerrors.NewCustomError(err, "an unknown error occurred during deleting a user")
	}

	return nil
}
