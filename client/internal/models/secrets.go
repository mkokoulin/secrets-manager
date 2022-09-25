// Package models includes client models
package models

import (
	"time"

	pb "github.com/mkokoulin/secrets-manager.git/client/internal/pb/secrets"
)

type Secret struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Type string `json:"type"`
	IsDeleted bool `json:"is_deleted"`
	Value map[string]string `json:"value"`
}

func (s *Secret) TransferValueData() []*pb.Data {
	var result []*pb.Data
	for k, v := range s.Value {
		result = append(result, &pb.Data{
			Title: k,
			Value: v,
		})
	}
	return result
}