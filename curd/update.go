package curd

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrjms/go-curd/v2/db"
)

func ApiUpdateHook[T any](before, after func(m *T) error, fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		err := ctx.ShouldBindJSON(&data)
		//写入ID
		id := ctx.MustGet("id")
		//value.Elem().FieldByName("id").Set(reflect.ValueOf(id))

		if err != nil {
			Error(ctx, err)
			return
		}

		if before != nil {
			if err := before(&data); err != nil {
				Error(ctx, err)
				return
			}
		}

		_, err = db.Engine.ID(id).Cols(fields...).Update(&data)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			if err := after(&data); err != nil {
				Error(ctx, err)
				return
			}
		}

		OK(ctx, &data)
	}
}
