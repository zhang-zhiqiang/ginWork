package conf

import (
	"baseframe/pkg/db"
	"baseframe/pkg/log"
	"time"
)

type Config struct {
	Server *Server
	Mysql  *db.Options
	Log    *log.Options
	Jwt    *JWT
}

type Server struct {
	Addr string
	Mode string
}

type JWT struct {
	Realm      string
	Key        string
	Timeout    time.Duration
	MaxRefresh time.Duration `mapstructure:"max-refresh"`
}
