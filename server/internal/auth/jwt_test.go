// Package auth includes everething for working with JWT
package auth

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	accessSecret := "secret"
	refreshSecret := "secret"

	type args struct {
		userID                     string
		accessTokenSecret          string
		refreshTokenSecret         string
		accessTokenLiveTimeMinutes int
		refreshTokenLiveTimeDays   int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Positive test",
			args: args{
				userID: "123",
				accessTokenSecret: accessSecret,
				refreshTokenSecret: refreshSecret,
				accessTokenLiveTimeMinutes: 10,
				refreshTokenLiveTimeDays: 10,
			},
			want: "wantToken",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken(tt.args.userID, tt.args.accessTokenSecret, tt.args.refreshTokenSecret, tt.args.accessTokenLiveTimeMinutes, tt.args.refreshTokenLiveTimeDays)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.AccessToken == "" {
				t.Errorf("CreateToken() = %v, want %v", got.AccessToken, tt.want)
			}
		})
	}
}
