package main

import (
	"baseframe/internal/goserver/conf"
	"baseframe/pkg/config"
	"baseframe/pkg/log"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultConfigName = "goserver.yaml"
)

var Config *conf.Config

func main() {
	config.LoadConfig("", defaultConfigName)

	Config = &conf.Config{}
	if err := viper.Unmarshal(Config); err != nil {
		panic("配置文件映射异常")
	}

	log.Init(Config.Log)
	log.Info(Config.Server.Addr)

	g := loadServer(Config)

	// set gin mode
	gin.SetMode(Config.Server.Mode)

	insecureServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", Config.Server.Addr),
		Handler: g,
	}

	go func() {
		if err := insecureServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Errorf("listen: %s\n", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof("Shutting down server...")

	if err := insecureServer.Shutdown(ctx); err != nil {
		log.Errorf("Insecure Server forced to shutdown:%v", err)

	}

	log.Infof("Server exiting")
}
