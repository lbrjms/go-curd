package db

import "github.com/lbrjms/go-curd/v2/config"

const MODULE = "database"

func init() {
	config.Register(MODULE, "type", "mysql")
	config.Register(MODULE, "url", "root:123qweLWJ@tcp(139.224.249.109:3306)/master?charset=utf8")
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "sync", true)
}
