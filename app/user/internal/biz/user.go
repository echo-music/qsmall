package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"qsmall/api/user"
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

	b.log.WithContext(ctx).Info("你好我好，大家好")

	_, err := b.repo.CreateUser(ctx, &User{ID: 10001})
	if ok := user.IsUserNotFound(err); ok {
		fmt.Println("用户不存在")
	}
	return nil, err
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
