package db

import (
	"github.com/lbrjms/go-curd/v2/config"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var Engine *xorm.Engine

func Open() error {
	var err error
	Engine, err = xorm.NewEngine(config.GetString(MODULE, "type"), config.GetString(MODULE, "url"))
	if err != nil {
		return err
	}
	if config.GetBool(MODULE, "debug") {
		Engine.ShowSQL(true)
		Engine.SetLogLevel(log.LOG_DEBUG)
	}
	if config.GetBool(MODULE, "sync") {
		Engine.Sync2(models...)
	}
	return nil
}
func Close() error {
	return Engine.Close()
}
