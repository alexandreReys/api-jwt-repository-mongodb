package validation

import (
	"encoding/json"
	"errors"

	errorHandler "github.com/alexandreReys/api-jwt-repository-mongodb/src/configuration/error_handler"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl ut.Translator
)

func init() {
	if val , ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *errorHandler.ErrorHandler {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors
	
	if errors.As(validation_err, &jsonErr) {
		return errorHandler.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []errorHandler.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := errorHandler.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}
			
			errorsCauses = append(errorsCauses, cause)
		}

		return errorHandler.NewBadRequestValidationError("Invalid fields", errorsCauses)
	} else {
		return errorHandler.NewBadRequestError("Error trying to convert fields")
	}	
}
