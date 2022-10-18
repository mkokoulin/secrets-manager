// Package models includes models for the server
package models

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestUser_MarshalJSON(t *testing.T) {
	id := uuid.New()

	type fields struct {
		ID        uuid.UUID
		Login     string
		Password  string
		CreatedAt time.Time
		IsDeleted bool
	}

	type want struct {
		ID string
		CreatedAt string
	}

	tests := []struct {
		name    string
		fields  fields
		want    want
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: fields{
				ID: id,
				Login: "string",
				Password: "string",
				IsDeleted: false,
			},
			want: want {
				ID: id.String(),
				CreatedAt: "0001-01-01T00:00:00Z",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:        tt.fields.ID,
				Login:     tt.fields.Login,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				IsDeleted: tt.fields.IsDeleted,
			}
			got, err := u.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			f := map[string]string{}

			err = json.Unmarshal(got, &f)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.want, want{ ID: f["id"], CreatedAt: f["created_at"] }) {
				t.Errorf("User.MarshalJSON() = %v, want %v", f["created_at"], tt.want)
			}
		})
	}
}
