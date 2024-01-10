package utils

import (
	"math"
	"strconv"
	"strings"
)

type StringValidator struct {
	required bool
	min      int
	max      int
}

func NewStringValidator(args []string) StringValidator {
	validator := StringValidator{}

	for i := 1; i < len(args); i++ {
		if args[i] == "required" {
			validator.required = true
		} else if strings.Contains(args[i], "min") {
			validator.min, _ = strconv.Atoi(args[i][4:])
		} else if strings.Contains(args[i], "max") {
			validator.max, _ = strconv.Atoi(args[i][4:])
		}
	}
	if validator.max == 0 {
		validator.max = math.MaxInt
	}

	return validator
}

func (sv StringValidator) Validate(obj interface{}, fieldName string) error {
	str := (obj).(*string)
	if sv.required && (str == nil || *str == "") {
		if str == nil {
			return &ValidatorNilError{fieldName: fieldName}
		} else if *str == "" {
			return &ValidatorEmptyStringError{fieldName: fieldName}
		}
	} else if str != nil && *str != "" {
		if sv.min >= len(*str) {
			return &ValidatorMinLengthStringError{fieldName: fieldName, minLength: sv.min}
		} else if sv.max <= len(*str) {
			return &ValidatorMaxLengthStringError{fieldName: fieldName, maxLength: sv.max}
		}
	}

	return nil
}
