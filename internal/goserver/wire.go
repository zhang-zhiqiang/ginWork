//go:build wireinject
// +build wireinject

package goserver

import (
	"baseframe/internal/goserver/conf"
	v1 "baseframe/internal/goserver/controller/v1"
	"baseframe/internal/goserver/data"
	"baseframe/internal/goserver/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func loadServer(c *conf.Config) *gin.Engine {
	panic(wire.Build(data.ProviderSet, service.ProviderSet, v1.ProviderSet, ProviderSet))
}
