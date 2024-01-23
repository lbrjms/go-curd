package main

import (
	"fmt"
	"github.com/lbrjms/go-curd/v2/api"
	"github.com/lbrjms/go-curd/v2/config"
	_ "github.com/lbrjms/go-curd/v2/docs"
	"github.com/lbrjms/go-curd/v2/web"
)

func main() {
	fmt.Println("=========")
	engine := web.CreateEngine()
	Startup(engine)
	engine.Serve()
}
func Startup(engine *web.Engine) error {
	err := config.Load()
	if err != nil {
		_ = config.Store()
	}
	//注册前端接口
	api.RegisterRoutes(engine.Group("/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(&engine.RouterGroup, "master")
	return nil
}
