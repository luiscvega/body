package body

import (
	"reflect"
	"strconv"
)

type Parser struct {
	params map[string][]string
	s      interface{}
	v      reflect.Value
	t      reflect.Type
}

func (p Parser) parse() (err error) {
	for i := 0; i < p.v.NumField(); i++ {
		var (
			field     = p.v.Field(i)
			tagName   = p.t.Field(i).Tag.Get("name")
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

	return
}

func Parse(params map[string][]string, s interface{}) (err error) {
	parser := Parser{
		params: params,
		s:      s,
		v:      reflect.ValueOf(s).Elem(),
		t:      reflect.ValueOf(s).Elem().Type()}

	err = parser.parse()

	return
}
