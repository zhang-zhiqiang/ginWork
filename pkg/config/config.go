package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var cfgFile string

func LoadConfig(cfg string, defaultName string) {
	viper.AutomaticEnv() // 读取匹配的环境变量
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc")
		rootPath, _ := os.Getwd()
		viper.AddConfigPath(filepath.Join(rootPath, "conf"))
		viper.SetConfigName(defaultName)
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "错误：无法读取配置文件(%s): %v\n", cfgFile, err)
		os.Exit(1)
	}
}
