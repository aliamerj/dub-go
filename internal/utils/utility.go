package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func BuildQueryURL(query *url.Values, options interface{}) {
	if query == nil {
		return
	}

	optValue := reflect.ValueOf(options)
	if optValue.Kind() == reflect.Ptr {
		optValue = optValue.Elem() // Dereference if a pointer is passed
	}
	optType := optValue.Type()

	for i := 0; i < optValue.NumField(); i++ {
		field := optValue.Field(i)
		fieldType := optType.Field(i)

		// Skip unexported fields
		if fieldType.PkgPath != "" {
			continue
		}

		paramName := toLowerCamel(fieldType.Name)

		switch field.Kind() {
		case reflect.Slice, reflect.Array:
			for j := 0; j < field.Len(); j++ {
				element := field.Index(j)
				if element.CanInterface() {
					switch element.Kind() {
					case reflect.String:
						str := element.String()
						if str != "" {
							query.Add(paramName, str)
						}
					}
				}
			}
		case reflect.String:
			if str := field.String(); str != "" {
				query.Add(paramName, str)
			}
		case reflect.Bool:
			if field.Bool() {
				query.Add(paramName, "true")
			}
		case reflect.Int:
			if intValue := field.Int(); intValue != 0 {
				query.Add(paramName, fmt.Sprintf("%d", intValue))
			}
		}
	}
}

func toLowerCamel(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	return strings.ToLower(string(r[0])) + string(r[1:])
}
