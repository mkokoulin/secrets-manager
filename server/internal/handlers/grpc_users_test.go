package handlers

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/mkokoulin/secrets-manager.git/server/internal/pb/users"
)

func TestGRPCUsers_CreateUser(t *testing.T) {
	type fields struct {
		UnimplementedUsersServer pb.UnimplementedUsersServer
		userService              UserServiceInterface
	}
	type args struct {
		ctx context.Context
		in  *pb.CreateUserRequiest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gu := &GRPCUsers{
				UnimplementedUsersServer: tt.fields.UnimplementedUsersServer,
				userService:              tt.fields.userService,
			}
			got, err := gu.CreateUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GRPCUsers.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GRPCUsers.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
