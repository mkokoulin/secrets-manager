// Package models includes models for the server
package models

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestNewRawSecretData(t *testing.T) {
	type args struct {
		secret Secret
	}

	value := map[string]string{
		"string": "string",
	}

	data, _ := json.Marshal(value)

	rsd := RawSecretData{
		UserID:    "123",
		ID:        "0c96fd62-a262-4133-85b9-eca10d13bc5b",
		Data:      data,
		Type:      "string",
		CreatedAt: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
	}

	rsd.Encrypt()

	tests := []struct {
		name    string
		args    args
		want    *RawSecretData
		wantErr bool
	}{
		{
			name: "Positive test",
			args: args{
				secret: Secret{
					UserID:   "123",
					SecretID: "123",
					Data: SecretData{
						CreatedAt: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
						Type:      "string",
						Value:     value,
					},
				},
			},
			want: &rsd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRawSecretData(tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRawSecretData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Data, tt.want.Data) {
				t.Errorf("NewRawSecretData() = %v, want %v", got, tt.want)
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
			fields: fields{
				ID:        "1",
				CreatedAt: time.Now(),
				UserID:    "2",
				Type:      "string",
				IsDeleted: false,
				Data:      []byte("123"),
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
			if err := rsd.Encrypt(); (err != nil) != tt.wantErr {
				t.Errorf("RawSecretData.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRawSecretData_Decrypt(t *testing.T) {
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
			fields: fields{
				ID:        "1",
				CreatedAt: time.Now(),
				UserID:    "2",
				Type:      "string",
				IsDeleted: false,
				Data:      []byte("123"),
			},
			wantErr: true,
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
			if err := rsd.Decrypt(); (err != nil) != tt.wantErr {
				t.Errorf("RawSecretData.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
