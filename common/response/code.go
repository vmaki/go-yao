package response

type ResCode int64

const (
	CodeSuccess         ResCode = 200
	CodeSysError        ResCode = 500
	CodeBadRequest      ResCode = 400
	CodeUnauthorized    ResCode = 401
	CodeNotFound        ResCode = 404
	CodeValidationErr   ResCode = 422
	CodeTooManyRequests ResCode = 429
)

// 自定义错误
const (
	CodeServerBusy ResCode = 10000 + iota
	CodeTokenExpired
	CodeTokenExpiredMaxRefresh
	CodeTokenMalformed
	CodeTokenInvalid
	CodeHeaderEmpty
	CodeHeaderMalformed
	CodeUserExist
	CodeUserNotExist
	CodeVerifyCodeErr
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "请求成功",
	CodeSysError:        "服务器内部错误，请稍后再试",
	CodeBadRequest:      "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式",
	CodeUnauthorized:    "身份验证失败，请稍后再试",
	CodeNotFound:        "路由未定义，请确认 url 和请求方法是否正确",
	CodeValidationErr:   "请求参数有误",
	CodeTooManyRequests: "请求过于频繁",

	CodeServerBusy:             "服务器繁忙",
	CodeTokenExpired:           "令牌已过期",
	CodeTokenExpiredMaxRefresh: "令牌已过最大刷新时间",
	CodeTokenMalformed:         "请求令牌格式有误",
	CodeTokenInvalid:           "请求令牌无效",
	CodeHeaderEmpty:            "需要认证才能访问！",
	CodeHeaderMalformed:        "请求头格式有误",
	CodeUserExist:              "用户已存在",
	CodeUserNotExist:           "用户不存在",
	CodeVerifyCodeErr:          "验证码错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}

	return msg
}
