package validator

type Validator interface {
	Validate(obj interface{}, fieldName string) error
}
