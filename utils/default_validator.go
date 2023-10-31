package utils

type DefaultValidator struct {

}

func NewDefaultValidator(args []string) DefaultValidator  {
	return DefaultValidator{}
}

func (dv DefaultValidator) Validate(obj interface{}, fieldName string) error {
	return nil
}

