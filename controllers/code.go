package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin
	CodeUsernameNull
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "成功响应",
	CodeInvalidParam:    "参数错误",
	CodeUserExist:       "用户存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeInvalidToken:    "无效的Token",
	CodeNeedLogin:       "需要登录",
	CodeUsernameNull:    "用户名为空",
}

func Msg(c ResCode) string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
