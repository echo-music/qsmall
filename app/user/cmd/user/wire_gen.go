// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"qsmall/app/user/internal/biz"
	"qsmall/app/user/internal/conf"
	"qsmall/app/user/internal/data"
	"qsmall/app/user/internal/server"
	"qsmall/app/user/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userBiz := biz.NewUserBiz(userRepo, logger)
	userService := service.NewUserService(userBiz)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}