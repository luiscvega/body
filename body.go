package body

import (
	"reflect"
	"strconv"
)

func Parse(params map[string][]string, s interface{}) (err error) {
	v := reflect.ValueOf(s).Elem()

	for i := 0; i < v.NumField(); i++ {
		field_value := v.Field(i)
		field := v.Type().Field(i)
		tagName := field.Tag.Get("name")

		param_val := params[tagName]

		if field_value.Kind() == reflect.String {
			field_value.SetString(param_val[0])
		}

		if field_value.Kind() == reflect.Int {
			integer, _ := strconv.ParseInt(param_val[0], 10, 64)
			field_value.SetInt(integer)
		}

		if field_value.Kind() == reflect.Slice {
			slice := reflect.MakeSlice(field_value.Type(), 0, 0)

			for _, s := range param_val {
				if field_value.Type().Elem().Kind() == reflect.String {
					slice = reflect.Append(slice, reflect.ValueOf(s))
				}

				if field_value.Type().Elem().Kind() == reflect.Int64 {
					integer, _ := strconv.ParseInt(s, 10, 64)
					slice = reflect.Append(slice, reflect.ValueOf(integer))
				}
			}

			field_value.Set(slice)
		}
	}

	return err
}
