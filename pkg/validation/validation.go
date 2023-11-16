package validation

import (
	"baseframe/pkg/log"
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
)

var validate *validator.Validate
var trans ut.Translator

func init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	if err := validate.RegisterValidation("phone", phone); err != nil {
		log.Warnf("参数校验 手机号规则注入失败 %v", err)
	}

	trans, _ = ut.New(zh.New()).GetTranslator("zh")

	_ = validate.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "{0}格式错误", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field(), fe.Field())
		return t
	})

	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Warnf("validation validate register default %v", err)
	}
}

func Check(data interface{}) error {
	if errs := validate.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}

	return nil
}

func phone(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^(13|14|15|17|18|19)[0-9]{9}$`, fl.Field().String())
	return ok
}
