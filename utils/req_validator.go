package utils

import (
	"reflect"
	"strings"
)


func Validate(obj interface{}) []error {
	var errors []error
	v := reflect.ValueOf(obj)

	for i := 0; i < v.NumField(); i++ {
		tag, ok := v.Type().Field(i).Tag.Lookup("validate")

		if !ok {
			continue
		}
		validator := getValidatorByField(tag)
		inter := v.Field(i).Interface()
		fieldName := v.Type().Field(i).Name
		err := validator.Validate(inter, fieldName)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func getValidatorByField(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "string":
		return NewStringValidator(args)
	}
	return NewDefaultValidator(args)
}