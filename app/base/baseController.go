package base

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

type BaseController struct {
}

type Response struct {
	Code    int
	Message string
	Result  interface{}
}

//回傳共用func
func (t *BaseController) ResponseJson(ctx *gin.Context, code int, message string, result interface{}) {

	if message == "" {
		message = http.StatusText(code)
	}

	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Result:  result,
	})

	ctx.Abort()
}

//檢核共用func
func (t *BaseController) ValidateRequest(request interface{}) (string, error) {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh_tw")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	err = validate.Struct(request)
	errMsg := ""
	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {

			if errMsg == "" {
				errMsg = err.Translate(trans)
			} else {
				errMsg = errMsg + " ," + err.Translate(trans)
			}
		}
	}

	return errMsg, err
}
