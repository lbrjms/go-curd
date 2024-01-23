package curd

import (
	"github.com/segmentio/ksuid"
	"reflect"
)

func GenerateKSUID[T any]() func(data *T) error {
	return func(data *T) error {
		value := reflect.ValueOf(data).Elem()
		field := value.FieldByName("id")
		//使用UUId作为Id
		//field.IsZero() 如果为空串时，生成UUID
		if field.Len() == 0 {
			key := ksuid.New().String()
			field.SetString(key)
		}
		return nil
	}
}
