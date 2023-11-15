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

func main() {
	config.LoadConfig("", defaultConfigName)

	options := &conf.Config{}
	if err := viper.Unmarshal(options); err != nil {
		panic("配置文件映射异常")
	}

	log.Init(options.Log)
	log.Info(options.Server.Addr)

	g := loadServer(options)

	// set gin mode
	gin.SetMode(options.Server.Mode)

	insecureServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", options.Server.Addr),
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
