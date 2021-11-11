package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPwd
	CodeServerBusy
	UnExceptionCode

	CodeInvalidToken
	CodeNeedLogin
)

var (
	errorMsg = map[ResCode]string{
		CodeSuccess:      "success",
		CodeInvalidParam: "无效的参数",
		CodeUserExist:    "用户已存在",
		CodeUserNotExist: "用户不存在",
		CodeInvalidPwd:   "无效的密码",
		CodeServerBusy:   "伺服器繁忙",
		UnExceptionCode:  "意外的情况",

		CodeInvalidToken: "无效的Token",
		CodeNeedLogin:    "需要登录",
	}
)

func (c ResCode) Msg() string {
	msg, ok := errorMsg[c]
	if !ok {
		return errorMsg[CodeServerBusy]
	}
	return msg
}
