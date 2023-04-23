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
	db    *db
	cache *cache
}

// NewData .
func NewData(c *conf.Data, db *db, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

type db struct {
}

//mysql实例
func NewDB() *db {
	return &db{}
}

type cache struct {
}

func NewCache() *cache {
	return &cache{}
}

func (s *userRepo) CreateUser(ctx context.Context, req *biz.User) (*biz.User, error) {
	//err := errors.New(404, "USER_NOT_FOUND", "user name is empty")

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
