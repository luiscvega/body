package body

import (
	"reflect"
	"strconv"
)

func Parse(params map[string][]string, s interface{}) (err error) {
	v := reflect.ValueOf(s).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		var (
			field     = v.Field(i)
			tagName   = t.Field(i).Tag.Get("name")
			param_val = params[tagName]
		)

		switch field.Kind() {
		case reflect.String:
			field.SetString(param_val[0])
		case reflect.Int:
			integer, _ := strconv.ParseInt(param_val[0], 10, 64)
			field.SetInt(integer)
		case reflect.Slice:
			slice := reflect.MakeSlice(field.Type(), 0, 0)

			for _, s := range param_val {
				if field.Type().Elem().Kind() == reflect.String {
					slice = reflect.Append(slice, reflect.ValueOf(s))
				}

				if field.Type().Elem().Kind() == reflect.Int64 {
					integer, _ := strconv.ParseInt(s, 10, 64)
					slice = reflect.Append(slice, reflect.ValueOf(integer))
				}
			}

			field.Set(slice)
		}
	}

	return err
}
