package utils_test

import (
	"errors"
	"go_auth_server/utils"
	"testing"
)

func TestUnrequieredNilStringValidation(t *testing.T) {
	s := struct {
		Str *string `validate:"string,min=5,max=100"`
	}{}

	errs := utils.Validate(s)

	if len(errs) > 0 {
		t.Fatalf("Unrequired nil string. errors: %v", errs)
	}
}

func TestUnrequiredEmptyStringValidation(t *testing.T) {
	empty_str := ""
	s := struct {
		Str *string `validate:"string,min=5,max=100"`
	}{&empty_str}

	errs := utils.Validate(s)

	if len(errs) > 0 {
		t.Fatalf("Unrequired empty string. errors: %v", errs)
	}
}

func TestRequiredNilStringValidation(t *testing.T) {
	s := struct {
		Str *string `validate:"string,required,min=5,max=100"`
	}{}

	errs := utils.Validate(s)

	var e *utils.ValidatorNilError
	contains := false
	for _, err := range errs {
		if errors.As(err, &e) {
			contains = true
		}
	}
	if !contains {
		t.Fatalf("Required nil string. errors: %v", errs)
	}
}

func TestRequiredEmptyStringValidation(t *testing.T) {
	empty_str := ""
	s := struct {
		Str *string `validate:"string,required,min=5,max=100"`
	}{&empty_str}

	errs := utils.Validate(s)

	var e *utils.ValidatorEmptyStringError
	contains := false
	for _, err := range errs {
		if errors.As(err, &e) {
			contains = true
		}
	}
	if !contains {
		t.Fatalf("Required empty string. errors: %v", errs)
	}
}

func TestUnrequiredMinLengthStringValidation(t *testing.T) {
	str := "123"
	s := struct {
		Str *string `validate:"string,min=5,max=100"`
	}{&str}

	errs := utils.Validate(s)

	var e *utils.ValidatorMinLengthStringError
	contains := false
	for _, err := range errs {
		if errors.As(err, &e) {
			contains = true
		}
	}
	if !contains {
		t.Fatalf("Unrequired non-empty string. String: %v, errors: %v", str, errs)
	}
}

func TestUnrequiredMaxLengthStringValidation(t *testing.T) {
	str := "123456"
	s := struct {
		Str *string `validate:"string,min=1,max=5"`
	}{&str}

	errs := utils.Validate(s)

	var e *utils.ValidatorMaxLengthStringError
	contains := false
	for _, err := range errs {
		if errors.As(err, &e) {
			contains = true
		}
	}
	if !contains {
		t.Fatalf("Unrequired non-empty string. String: %v, errors: %v", str, errs)
	}
}
