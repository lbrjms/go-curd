package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrjms/go-curd/v2/curd"
	"github.com/lbrjms/go-curd/v2/types"
)

// @Summary 查询用户
// @Schemes
// @Description 查询用户
// @Tags user
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.User] 返回用户信息
// @Router /user/list [get]
func noopUserList() {}

// @Summary 创建用户
// @Schemes
// @Description 创建用户
// @Tags user
// @Param search body types.User true "用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.User] 返回用户信息
// @Router /user/create [post]
func noopUserCreate() {}

// @Summary 修改用户
// @Schemes
// @Description 修改用户
// @Tags user
// @Param id path string true "用户ID"
// @Param user body types.User true "用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.User] 返回用户信息
// @Router /user/{id} [post]
func noopUserUpdate() {}

// @Summary 禁用用户
// @Schemes
// @Description 禁用用户
// @Tags user
// @Param id path string true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.User] 返回用户信息
// @Router /user/{id}/disable [get]
func noopUserDisable() {}

func UserRouter(app *gin.RouterGroup) {
	app.POST("/create", curd.ApiCreateHook[types.User](curd.GenerateKSUID[types.User](), nil))
	app.GET("/:id/disable", curd.ParseParamStringId, curd.ApiDisableHook[types.User](true, nil, nil))
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.User](nil, nil,
		"name", "cellphone", "email", "roles", "disabled"))

	app.GET("/list", curd.AipList[types.User]())
}
