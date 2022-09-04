package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
	"gorm.io/gorm"
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
	tx := pd.conn.Create(user)

	if tx.Error != nil {
		return user.ID, tx.Error
	}

	return user.ID, nil
}

func (pd *PostgresDatabase) CheckUserPassword(ctx context.Context, user models.User) (string, error) {
	return "", nil
}

func (pd *PostgresDatabase) DeleteUser(ctx context.Context, userID uuid.UUID) error {	
	tx := pd.conn.Model(&models.User{}).Where("id=?", userID.String()).Update("is_deleted", true)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}



