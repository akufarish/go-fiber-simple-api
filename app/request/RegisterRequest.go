package request

import "gopkg.in/go-playground/validator.v9"

type(
	RegisterRequest struct {
		Email    string `validate:"required,min=8,email"`
		Username string `validate:"required,min=10"`
		Password string `validate:"required"`
	}

	Validator struct {
		validator *validator.Validate
	}

	ErrorResponse struct {
        Error       bool
        FailedField string
        Tag         string
        Value       interface{}
    }
)

var validate = validator.New()

func (v Validator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	err := validate.Struct(data) 

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}			        