package goserver

import (
	"baseframe/internal/goserver/controller/v1/user"
	"baseframe/pkg/core"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewHttpServer)

func NewHttpServer(uc *user.UserController) *gin.Engine {
	g := gin.New()
	g.GET("/ping", func(c *gin.Context) {
		core.Response(c, nil, map[string]string{"message": "pong"})
	})

	users := g.Group("/users")
	{
		users.POST("register", uc.Create)
	}
	return g
}
