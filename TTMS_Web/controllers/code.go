package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedAuth
	CodeInvalidToken
	CodeNeedLogin
	CodeConflictSchedule
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "无效的参数",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "无效密码",
	CodeServerBusy:      "系统繁忙",

	CodeNeedAuth:     "需要Auth",
	CodeInvalidToken: "无效的Token",
	CodeNeedLogin:    "需要登录",
	CodeConflictSchedule: "演出计划冲突",
}

func (c ResCode) Msg() string {
	v, ok := codeMsgMap[c]
	if !ok {
		v = codeMsgMap[CodeServerBusy]
	}
	return v
}