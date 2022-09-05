package database

import (
	"context"

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

func (pd *PostgresDatabase) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
        return user.ID, err
    }

	user.Password = string(hash)
	
	tx := pd.conn.Create(user)

	if tx.Error != nil {
		return user.ID, tx.Error
	}

	return user.ID, nil
}

func (pd *PostgresDatabase) CheckUserPassword(ctx context.Context, user models.User) (uuid.UUID, error) {
	var result models.User

	tx := pd.conn.Model(&models.User{}).Where("login = ?", user.Login).First(result)

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
        return user.ID, err
    }

	if tx.Error != nil {
		return user.ID, tx.Error
	}

	return user.ID, nil
}

func (pd *PostgresDatabase) DeleteUser(ctx context.Context, userID uuid.UUID) error {	
	tx := pd.conn.Model(&models.User{}).Where("id=?", userID.String()).Update("is_deleted", true)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}



