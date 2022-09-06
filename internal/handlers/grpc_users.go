package handlers

import pb "github.com/mkokoulin/secrets-manager.git/internal/pb/users"

type GRPCUsers struct {
	pb.UnimplementedUsersServer
	userService UserServiceInterface
}