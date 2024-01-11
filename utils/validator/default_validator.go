package validator

type DefaultValidator struct {
}

func NewDefaultValidator() DefaultValidator {
	return DefaultValidator{}
}

func (dv DefaultValidator) Validate(obj interface{}, fieldName string) error {
	return nil
}
