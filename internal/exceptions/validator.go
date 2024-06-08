package exceptions

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

var (
	validate = validator.New()
)

func ValidateStruct(obj interface{}) error {
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	sf := strings.ToLower(validationError.StructField())
	pm := validationError.Param()
	switch validationError.Tag() {
	case "required":
		return errors.New(sf + " is required")
	case "max":
		return errors.New(sf + " reached max of " + pm + " characters")
	case "min":
		return errors.New(sf + " must have " + pm + " characters at least")
	case "email":
		return errors.New(sf + " is not valid")
	}

	return nil
}
