package goserver

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

	"github.com/spf13/cobra"
)

const (
	defaultConfigName = "goserver.yaml"
)

var Config *conf.Config

func NewGoServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "goserver",
		Short:        "gin 基础项目",
		Long:         ``,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		PostRun: func(cmd *cobra.Command, args []string) {},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}
	return cmd
}

func run() error {
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

		return err
	}

	log.Infof("Server exiting")

	return nil
}
