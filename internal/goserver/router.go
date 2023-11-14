package goserver

import (
	"baseframe/internal/goserver/controller/v1/user"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var ProviderSet = wire.NewSet(NewHttpServer)

func NewHttpServer(uc *user.UserController) *gin.Engine {
	g := gin.New()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong2134",
		})
	})

	users := g.Group("/users")
	{
		users.GET("create", uc.Create)
	}
	return g
}
