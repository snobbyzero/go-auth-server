package validator

import "fmt"

type ValidatorNilError struct{ fieldName string }

func (e *ValidatorNilError) Error() string {
	return fmt.Sprintf("%v shouldn't be null", e.fieldName)
}

type ValidatorEmptyStringError struct{ fieldName string }

func (e *ValidatorEmptyStringError) Error() string {
	return fmt.Sprintf("%v shouldn't be empty", e.fieldName)
}

type ValidatorMinLengthStringError struct {
	fieldName string
	minLength int
}

func (e *ValidatorMinLengthStringError) Error() string {
	return fmt.Sprintf("%v should be more than %v", e.fieldName, e.minLength)
}

type ValidatorMaxLengthStringError struct {
	fieldName string
	maxLength int
}

func (e *ValidatorMaxLengthStringError) Error() string {
	return fmt.Sprintf("%v should be less than %v", e.fieldName, e.maxLength)
}
