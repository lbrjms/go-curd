package curd

import "github.com/gin-gonic/gin"

func AipList[T any](fields ...string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var body ParamList
		err := context.ShouldBindQuery(&body)
		if err != nil {
			Error(context, err)
		}

		query := body.ToQuery()
		if len(fields) > 0 {
			query.Cols(fields...)
		}
		var datum []T
		count, err := query.FindAndCount(&datum)
		if err != nil {
			Error(context, err)
			return
		}
		List(context, datum, count)
	}
}
