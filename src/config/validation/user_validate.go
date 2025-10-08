package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	pt "github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	resterr "github.com/rlevidev/oauth-go/src/config/rest_err"

	pt_translation "github.com/go-playground/validator/v10/translations/pt"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt := pt.New()
		unicode := ut.New(pt, pt)

		transl, _ = unicode.GetTranslator("pt")
		pt_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *resterr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError("Campo inválido")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []resterr.Causes{}

		for _, err := range validation_err.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: err.Translate(transl),
				Field:   err.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		return resterr.NewBadRequestValidationError("Erro de validação", errorsCauses)
	} else {
		return resterr.NewInternalServerError("Erro ao processar a requisição")
	}
}
