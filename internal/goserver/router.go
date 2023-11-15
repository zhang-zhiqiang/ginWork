package goserver

import (
	"baseframe/internal/goserver/controller/v1/user"
	"baseframe/internal/pkg/code"
	"baseframe/pkg/core"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewHttpServer)

func NewHttpServer(uc *user.UserController) *gin.Engine {
	g := gin.New()
	g.GET("/ping", func(c *gin.Context) {
		data := &gin.H{"message": "pong"}
		core.WriteResponse(c, code.ErrSuccess, data)
	})

	users := g.Group("/users")
	{
		users.GET("create", uc.Create)
	}
	return g
}
