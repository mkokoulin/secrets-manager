package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

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

func (pd *PostgresDatabase) CreateUser(ctx context.Context, user models.User) (string, error) {
	err := pd.conn.Transaction(func(tx *gorm.DB) error {
		var exists bool

		err := tx.Model(&models.User{}).Select("count(*) > 0").Where("login = ?", user.Login).Find(&exists).Error
		if err != nil {
			return err
		}

		if exists {
			return errors.New("user already exists")
		}


		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user.Password = string(hash)

		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return user.ID.String(), nil
}

func (pd *PostgresDatabase) CheckUserPassword(ctx context.Context, login, password string) (string, error) {
	var result models.User

	err := pd.conn.Model(&models.User{}).Where("login = ?", login).First(&result).Error
	if err != nil {
		return result.Login, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
        return result.Login, err
    }

	return result.Login, nil
}

func (pd *PostgresDatabase) DeleteUser(ctx context.Context, userID uuid.UUID) error {	
	err := pd.conn.Model(&models.User{}).Where("id=?", userID.String()).Update("is_deleted", true).Error
	if err != nil {
		return err
	}

	return nil
}



