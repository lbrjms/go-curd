package web

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/lbrjms/go-curd/v2/config"
	"strconv"
)

type Engine struct {
	*gin.Engine
}

func CreateEngine() *Engine {
	if !config.GetBool(MODULE, "debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()
	app.Use(gin.Recovery())
	if config.GetBool(MODULE, "debug") {
		app.Use(gin.Logger())
	}
	if config.GetBool(MODULE, "cors") {
		c := cors.DefaultConfig()
		c.AllowAllOrigins = true
		c.AllowCredentials = true
		app.Use(cors.New(c))
	}
	app.Use(sessions.Sessions("gin-test", cookie.NewStore([]byte("gin-test"))))
	if config.GetBool(MODULE, "gzip") {
		app.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	return &Engine{app}
}

func (app *Engine) Serve() {
	port := config.GetInt(MODULE, "port")
	addr := ":" + strconv.Itoa(port)

	err := app.Run(addr)
	if err != nil {
		fmt.Println(err)
	}
}
