package body

import (
	"reflect"
	"strconv"
)

func Parse(params map[string][]string, s interface{}) (err error) {
	val := reflect.ValueOf(s).Elem()
	typ := reflect.ValueOf(s).Elem().Type()

	for i := 0; i < val.NumField(); i++ {
		var (
			field    = val.Field(i)
			tagName  = typ.Field(i).Tag.Get("name")
			paramVal []string
			match    bool
		)

		if paramVal, match = params[tagName]; match == false {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(paramVal[0])
		case reflect.Int:
			integer, _ := strconv.ParseInt(paramVal[0], 10, 64)
			field.SetInt(integer)
		case reflect.Slice:
			slice := reflect.MakeSlice(field.Type(), 0, 0)

			for _, item := range paramVal {
				switch field.Type().Elem().Kind() {
				case reflect.String:
					slice = reflect.Append(slice, reflect.ValueOf(item))
				case reflect.Int64:
					integer, _ := strconv.ParseInt(item, 10, 64)
					slice = reflect.Append(slice, reflect.ValueOf(integer))
				}
			}

			field.Set(slice)
		}
	}

	return
}
