package objconv

import (
	"reflect"
	"strings"
)

func contains(arr []string, obj string) bool {
	for _, ele := range arr {
		if ele == obj {
			return true
		}
	}
	return false
}

func ToMap(object interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	target := reflect.ValueOf(object)
	elements := target.Elem()

	for i := 0; i < elements.NumField(); i++ {
		val := elements.Field(i)
		mType := elements.Type().Field(i)
		jsonTag := mType.Tag.Get("json")

		if jsonTag != "-" && jsonTag != "" {
			tagE := strings.Split(jsonTag, ",")
			if len(tagE) == 0 {
				continue
			}
			m[tagE[0]] = val.Interface()
		}
	}
	return m
}

func ToMapWithFields(object interface{}, fields []string) map[string]interface{} {
	if len(fields) == 0 {
		return ToMap(object)
	}

	m := make(map[string]interface{})

	target := reflect.ValueOf(object)
	elements := target.Elem()

	for i := 0; i < elements.NumField(); i++ {
		val := elements.Field(i)
		mType := elements.Type().Field(i)
		jsonTag := mType.Tag.Get("json")

		if jsonTag != "-" && jsonTag != "" {
			tagE := strings.Split(jsonTag, ",")
			if len(tagE) == 0 || !contains(fields, tagE[0]) {
				continue
			}
			m[tagE[0]] = val.Interface()
		}
	}
	return m
}
