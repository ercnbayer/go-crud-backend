package validator

import (
	"go-backend/logger"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateID(id string) error {

	return Validate.Var(id, "uuid4")
}

func ValidateUpdatedStruct(person interface{}) error {

	logger.Info("validating partial")

	return Validate.StructExcept(person)

}

func ValidateStruct(person interface{}) error {

	logger.Info("validating normal")
	return Validate.Struct(person)
}
