package curd

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrjms/go-curd/v2/db"
	"xorm.io/xorm"
)

type ParamList struct {
	Skip  int `form:"skip" json:"skip"`
	Limit int `form:"limit" json:"limit"`
}
type ParamStringId struct {
	Id string `uri:"id"`
}

func (body *ParamList) ToQuery() *xorm.Session {
	if body.Limit < 1 {
		body.Limit = 20
	}
	op := db.Engine.Limit(body.Limit, body.Skip)
	op.Desc("id")
	return op
}

func ParseParamStringId(ctx *gin.Context) {
	var pid ParamStringId
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		Error(ctx, err)
		ctx.Abort()
		return
	}
	ctx.Set("id", pid.Id)
	ctx.Next()
}
