package handlers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"github.com/mkokoulin/secrets-manager.git/server/internal/models"
	pb "github.com/mkokoulin/secrets-manager.git/server/internal/pb/users"
)

func TestGRPCUsers_CreateUser(t *testing.T) {
	type args struct {
		in  *pb.CreateUserRequiest
	}
	tests := []struct {
		name    string
		args    args
		mockError error
		want    *pb.CreateUserResponse
		wantErr bool
	}{
		{
			name: "Test #1",
			mockError: nil,
			args: args {
				in: &pb.CreateUserRequiest {
					Login: "",
					Password: "",
				},
			},
		},	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := models.User {
				ID: uuid.New(),
				Login: tt.args.in.Login,
				Password: tt.args.in.Password,
			}

			userServiceMock := NewMockUserServiceInterface(ctrl)

			userServiceMock.EXPECT().CreateUser(gomock.Any(), user).Return(tt.mockError).AnyTimes()
		})
	}
}
