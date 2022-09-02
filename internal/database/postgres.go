package database

import (
	"context"

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

func (pd *PostgresDatabase) CreateUser(ctx context.Context, user models.User) error {
	tx := pd.conn.Create(user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (pd *PostgresDatabase) CheckUserPassword(ctx context.Context, user models.User) (string, error) {
	
}

func (pd *PostgresDatabase) DeleteUser(ctx context.Context, userID string) error {

}



