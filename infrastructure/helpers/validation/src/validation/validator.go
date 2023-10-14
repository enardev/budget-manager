package validation

import (
	"encoding/json"
	"io"

	"github.com/enaldo1709/budget-manager/infrastructure/helpers/validation/src/validation/errors"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New()
	return &Validator{validate: validate}
}

func (v *Validator) DecodeAndValidate(r io.Reader, s interface{}) error {
	if err := json.NewDecoder(r).Decode(s); err != nil {
		return errors.NewDecodeError(s, err)
	}
	if err := v.validate.Struct(s); err != nil {
		return errors.NewValidationError(s, err)
	}
	return nil
}
