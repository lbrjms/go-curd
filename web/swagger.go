package web

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwaggerDocs(app *gin.RouterGroup, instance string) {
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
