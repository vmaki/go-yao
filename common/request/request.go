package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"go-yao/common/response"
)

func Validate(ctx *gin.Context, req IRequest) error {
	if err := ctx.ShouldBind(req); err != nil {
		return response.New(response.CodeBadRequest)
	}

	// 表单验证
	err := req.Generate(req)
	if err != nil {
		return err
	}

	return nil
}

func GoValidate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) error {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	errs := govalidator.New(opts).ValidateStruct()
	if len(errs) > 0 {
		str := ""
		for _, v := range errs {
			str = v[0]
			break
		}

		return response.New(response.CodeValidationErr, str)
	}

	return nil
}
