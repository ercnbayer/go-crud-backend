package validator

import (
	"go-backend/logger"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateID(id string) error {

	return Validate.Var(id, "uuid4")
}

// You can use persons pointer
func ValidateUpdatedStruct(person interface{}) error {

	logger.Info("validating partial")

	return Validate.StructExcept(person)

}

//Also you need to map inside validatione erros like
//for _, err := range err.(validator.ValidationErrors)
//After that you can add this errors to your validation error DTO like that
/*
var element customerror.ValidationErrorData
element.FailedField = err.StructNamespace()
element.Tag = err.Tag()
element.Value = err.Param()
*/
//After finish this loop you can create and return your validation error. And you can control this error type and return is from your customer fiber error handler https://docs.gofiber.io/guide/error-handling/
func ValidateStruct(person interface{}) error {

	logger.Info("validating normal")
	return Validate.Struct(person)
}
