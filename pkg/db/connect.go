package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/1005281342/task-manager/internal/config"
)

type Connection struct {
	db *gorm.DB
}

func (c Connection) GetDB() *gorm.DB {
	return c.db
}

func New(cfg config.Config) Connection {
	var d gorm.Dialector
	switch cfg.Gorm.Driver {
	case "sqlite":
		d = sqlite.Open(cfg.Gorm.Dsn)
	case "postgres":
		d = postgres.Open(cfg.Gorm.Dsn)
	default:
		panic(fmt.Errorf("unsupported driver:" + cfg.Gorm.Driver))
	}

	var db, err = gorm.Open(d)
	if err != nil {
		panic(err)
	}
	return Connection{db: db}
}
