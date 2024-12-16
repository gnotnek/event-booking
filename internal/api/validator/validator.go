package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

// NewValidator creates a new instance of the Validator service.
func NewValidator() *Validator {
	validate := validator.New()

	// Register custom regex validation for names
	_ = validate.RegisterValidation("name", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$`)
		return re.MatchString(fl.Field().String())
	})

	return &Validator{validate: validate}
}

// ValidateStruct validates a struct based on tags
func (v *Validator) ValidateStruct(payload interface{}) error {
	err := v.validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%s is invalid: %s", err.Field(), err.Tag())
		}
	}
	return nil
}
