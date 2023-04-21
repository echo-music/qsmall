package service

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"qsmall/app/user/internal/biz"

	pb "qsmall/api/user"
)

type UserService struct {
	pb.UnimplementedUserServer
	bz *biz.UserBiz
}

func NewUserService(bz *biz.UserBiz) *UserService {
	return &UserService{bz: bz}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	ctx, span := otel.Tracer("user").Start(ctx, "CreateUser")
	defer span.End()
	span.SetAttributes(attribute.String("key1", "value"), attribute.Int("key2", 123))

	_, err := s.bz.CreateUser(ctx, &biz.User{
		ID: 10001,
	})
	fmt.Println("hhecevcervervre")
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
