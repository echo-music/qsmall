package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"qsmall/app/user/internal/biz"
	"qsmall/app/user/internal/conf"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type Data struct {
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func (s *userRepo) CreateUser(ctx context.Context, req *biz.User) (*biz.User, error) {
	return nil, nil
}
func (s *userRepo) UpdateUser(ctx context.Context, req *biz.User) (*biz.User, error) {
	return nil, nil
}
func (s *userRepo) DeleteUser(ctx context.Context, req *biz.User) (*biz.User, error) {
	return nil, nil
}
func (s *userRepo) GetUser(ctx context.Context, req *biz.User) (*biz.User, error) {
	return nil, nil
}
func (s *userRepo) ListUser(ctx context.Context, req *biz.User) (*biz.User, error) {
	return nil, nil
}
