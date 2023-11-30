package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var Validator = validator.New()

func ValidateID(id string) (uuid.UUID, error) {

	return uuid.Parse(id)
}
