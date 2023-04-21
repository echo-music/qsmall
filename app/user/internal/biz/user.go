package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID int `json:"id,omitempty"`
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, *User) (*User, error)
	GetUser(context.Context, *User) (*User, error)
	ListUser(context.Context, *User) (*User, error)
}

type UserBiz struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserBiz(repo UserRepo, logger log.Logger) *UserBiz {
	return &UserBiz{repo: repo, log: log.NewHelper(logger)}
}

func (b *UserBiz) CreateUser(ctx context.Context, req *User) (*User, error) {
	return b.repo.CreateUser(ctx, &User{ID: 10001})
}
func (b *UserBiz) UpdateUser(ctx context.Context, req *User) (*User, error) {
	return nil, nil
}
func (b *UserBiz) DeleteUser(ctx context.Context, req *User) (*User, error) {
	return nil, nil
}
func (b *UserBiz) GetUser(ctx context.Context, req *User) (*User, error) {
	return nil, nil
}
func (b *UserBiz) JJ(ctx context.Context, req *User) (*User, error) {
	return nil, nil
}
