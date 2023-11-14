package data

import (
	"baseframe/internal/goserver/conf"
	"baseframe/pkg/db"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
}

var ProviderSet = wire.NewSet(NewData, NewUserData)

// NewData .
func NewData(c *conf.Config) *Data {
	db, err := db.New(c.Mysql)
	if err != nil {
		panic("数据库连接失败")
	}
	return &Data{db: db}
}
