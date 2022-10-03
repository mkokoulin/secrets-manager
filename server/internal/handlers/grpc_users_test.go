package handlers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"github.com/mkokoulin/secrets-manager.git/client/internal/models"
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

			// gu := &GRPCUsers{
			// 	UnimplementedUsersServer: tt.fields.UnimplementedUsersServer,
			// 	userService:              tt.fields.userService,
			// }
			// got, err := gu.CreateUser(tt.args.ctx, tt.args.in)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("GRPCUsers.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GRPCUsers.CreateUser() = %v, want %v", got, tt.want)
			// }
		})
	}
}
