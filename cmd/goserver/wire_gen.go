// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"baseframe/internal/goserver"
	"baseframe/internal/goserver/conf"
	"baseframe/internal/goserver/controller/v1/user"
	"baseframe/internal/goserver/data"
	"baseframe/internal/goserver/service/v1"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func loadServer(c *conf.Config) *gin.Engine {
	db := data.NewData(c)
	userRepo := data.NewUserData(db)
	userService := v1.NewUserService(userRepo)
	userController := user.NewUserController(userService)
	ginJWTMiddleware := user.NewJwtAuth(c, userController)
	engine := goserver.NewHttpServer(userController, ginJWTMiddleware)
	return engine
}
