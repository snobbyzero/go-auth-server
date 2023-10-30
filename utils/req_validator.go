package utils

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type Validator interface {
	Validate(obj interface{}, fieldName string) error
}

type DefaultValidator struct {

}

func (dv DefaultValidator) Validate(obj interface{}, fieldName string) error {
	return nil
}

type StringValidator struct {
	required bool
	min int
	max int
}

func (sv StringValidator) Validate(obj interface{}, fieldName string) error {
	str := (obj).(*string)
	if sv.required && (str == nil || *str == "") {
		return fmt.Errorf("%v shouldn't be empty or null", fieldName)
	} else if sv.min > len(*str) {
		return fmt.Errorf("%v should be more than %v", fieldName, sv.min)
	} else if sv.max < len(*str) {
		return fmt.Errorf("%v should be less than %v", fieldName, sv.max)
	}

	return nil
}

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
		validator := StringValidator{}
		if strings.Contains(tag, "required") {
			validator.required = true
		}

		if _, err := fmt.Sscanf(tag, "min=%d", &validator.min); err != nil {
			validator.min = 0
		}
		if _, err := fmt.Sscanf(tag, "max=%d", &validator.max); err != nil {
			validator.max = math.MaxInt
		}
		return validator
	}
	return DefaultValidator{}
}