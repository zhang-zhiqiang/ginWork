package conf

import (
	"baseframe/pkg/db"
	"baseframe/pkg/log"
)

type Server struct {
	Addr string
	Mode string
}

type Config struct {
	Server *Server
	Mysql  *db.Options
	Log    *log.Options
}
