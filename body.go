package body

import (
	"fmt"
	"reflect"
	"strconv"
)

func Parse(params map[string][]string, s interface{}) (err error) {
	parser := Parser{
		params: params,
		struc:  s,
		val:    reflect.ValueOf(s).Elem(),
		typ:    reflect.ValueOf(s).Elem().Type()}

	err = parser.parse()

	return
}

type Parser struct {
	params map[string][]string
	struc  interface{}
	val    reflect.Value
	typ    reflect.Type
}

func (p Parser) parse() (err error) {
	for i := 0; i < p.val.NumField(); i++ {
		var (
			field     = p.val.Field(i)
			tagName   = p.typ.Field(i).Tag.Get("name")
			param_val = p.params[tagName]
		)

		switch field.Kind() {
		case reflect.String:
			field.SetString(param_val[0])
		case reflect.Int:
			integer, _ := strconv.ParseInt(param_val[0], 10, 64)
			field.SetInt(integer)
		case reflect.Slice:
			slice := reflect.MakeSlice(field.Type(), 0, 0)

			for _, item := range param_val {
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
