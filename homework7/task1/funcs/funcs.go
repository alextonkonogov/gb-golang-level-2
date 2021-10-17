package funcs

import (
	"fmt"
	"reflect"

	"github.com/alextonkonogov/gb-golang-level-2/homework7/task1/persons"
)

func ChangeStructField(in *persons.Person, values map[string]interface{}) (err error) {
	getAttr := func(obj interface{}, fieldName string) (reflect.Value, error) {
		pointToStruct := reflect.ValueOf(obj)
		curStruct := pointToStruct.Elem()
		curField := curStruct.FieldByName(fieldName)
		if !curField.IsValid() {
			return reflect.Value{}, fmt.Errorf("not found: %s", fieldName)
		}
		return curField, nil
	}

	for key, val := range values {
		switch v := val.(type) {
		case int:
			attr, err := getAttr(in, key)
			if err != nil {
				return err
			}
			attr.SetInt(int64(v))
		case string:
			attr, err := getAttr(in, key)
			if err != nil {
				return err
			}
			attr.SetString(v)
		case bool:
			attr, err := getAttr(in, key)
			if err != nil {
				return err
			}
			attr.SetBool(v)
		case float64:
			attr, err := getAttr(in, key)
			if err != nil {
				return err
			}
			attr.SetFloat(v)
		default:
			return fmt.Errorf("not found: %s", key)
		}
	}
	return
}
