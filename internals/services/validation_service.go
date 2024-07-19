package services

import (
	"reflect"
	"strings"
)

func GetEmptyField(body any) (bool, string) {
	v := reflect.ValueOf(body)
	t := reflect.TypeOf(body)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		if field.Kind() == reflect.String && field.String() == "" {
			return false, strings.ToLower(fieldName) 
		}
	}

	return true, ""
}