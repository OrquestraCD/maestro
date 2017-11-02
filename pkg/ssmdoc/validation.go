package ssmdoc

import (
	"fmt"
	"reflect"
)

func invalidField(fieldName string) error {
	return fmt.Errorf("field %s must be set", fieldName)
}

func validateStep(step interface{}) error {
	val := reflect.ValueOf(step)

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("Invalid plugin type %s, expected Struct", val.Kind())
	}

	structTyp := reflect.TypeOf(step)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := structTyp.Field(i)
		tag := fieldType.Tag.Get("required")

		if tag == "true" {
			switch field.Kind() {
			case reflect.Int:
				if field.Int() == 0 {
					return invalidField(fieldType.Name)
				}
			case reflect.String:
				fallthrough
			case reflect.Slice:
				fallthrough
			case reflect.Array:
				if field.Len() == 0 {
					return invalidField(fieldType.Name)
				}
			case reflect.Struct:
				if err := validateStep(field.Interface()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
