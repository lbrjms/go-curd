package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	UserRouter(group.Group("/user"))
}
