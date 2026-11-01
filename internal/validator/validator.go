package validator

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type ValidatorErrors map[string]string

type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
	locale     string
}

func NewValidator() *Validator {
	idLocale := id.New()
	uni := ut.New(idLocale, idLocale)
	trans, found := uni.GetTranslator("id")
	if !found {
		log.Fatal("failed to get translator")
	}

	validate := validator.New()

	if err := id_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalf("failed to register translations: %v", err)
	}

	validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// misal buat overide translation message
	validate.RegisterTranslation("required", trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} tidak boleh kosong", true) // true = override
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)

	return &Validator{
		validate:   validate,
		translator: trans,
		locale:     "id",
	}
}

func (v *Validator) ValidateStruct(data any) ValidatorErrors {
	if err := v.validate.Struct(data); err != nil {
		return v.formatErrors(err)
	}
	return nil
}

func (v *Validator) formatErrors(err error) ValidatorErrors {
	out := ValidatorErrors{}
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range errs {
			out[fe.Field()] = fe.Translate(v.translator)
		}
	}
	return out
}

func (v *Validator) GetLocale() string {
	return v.locale
}
