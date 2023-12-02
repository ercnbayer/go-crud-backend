package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateID(id string) error {

	return validate.Var(id, "uuid4")
}

func ValidateUpdatedStruct(person interface{}) error {

	return validate.StructPartial(person)
}

func ValidateStruct(person interface{}) error {

	return validate.Struct(person)
}
