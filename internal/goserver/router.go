package goserver

import (
	"baseframe/internal/goserver/controller/v1/user"
	"baseframe/internal/pkg/code"
	"baseframe/pkg/core"
	"baseframe/pkg/errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewHttpServer)

func NewHttpServer(uc *user.UserController, auth *jwt.GinJWTMiddleware) *gin.Engine {
	g := gin.New()

	g.NoRoute(func(c *gin.Context) {
		core.Response(c, errors.WithCode(code.ErrPageNotFound, ""), nil)
	})

	g.GET("/ping", func(c *gin.Context) {
		core.Response(c, nil, map[string]string{"message": "pong"})
	})

	g.POST("/login", auth.LoginHandler)
	g.POST("/logout", auth.LogoutHandler)
	g.POST("/refresh", auth.RefreshHandler)

	g.Use(auth.MiddlewareFunc())
	users := g.Group("/users")
	{
		users.POST("register", uc.Create)
	}
	return g
}
