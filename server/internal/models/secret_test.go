// Package models includes models for the server
package models

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestRawSecretData_MarshalJSON(t *testing.T) {
	type fields struct {
		ID        string
		CreatedAt time.Time
		UserID    string
		Type      string
		IsDeleted bool
		Data      []byte
	}

	type want struct {
		CreatedAt string
	}

	tests := []struct {
		name    string
		fields  fields
		want    want
		wantErr bool
	}{
		{
			name:   "Test 1",
			fields: fields{},
			want: want{
				CreatedAt: "0001-01-01T00:00:00Z",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsd := &RawSecretData{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UserID:    tt.fields.UserID,
				Type:      tt.fields.Type,
				IsDeleted: tt.fields.IsDeleted,
				Data:      tt.fields.Data,
			}
			got, err := rsd.MarshalJSON()
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

			if !reflect.DeepEqual(tt.want, want{CreatedAt: f["created_at"]}) {
				t.Errorf("User.MarshalJSON() = %v, want %v", f["created_at"], tt.want)
			}
		})
	}
}

func TestRawSecretData_Encrypt(t *testing.T) {
	type fields struct {
		ID        string
		CreatedAt time.Time
		UserID    string
		Type      string
		IsDeleted bool
		Data      []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test #1",
			fields: fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsd := &RawSecretData{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UserID:    tt.fields.UserID,
				Type:      tt.fields.Type,
				IsDeleted: tt.fields.IsDeleted,
				Data:      tt.fields.Data,
			}

			err := rsd.Encrypt()
			if err != nil {
				t.Errorf("Secret.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var r []byte


			if rsd.Data == r {
				t.Errorf("RawSecretData.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
