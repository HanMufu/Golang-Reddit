package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "wrong parameters",
	CodeUserExist:       "User already existed",
	CodeUserNotExist:    "User does not exist",
	CodeInvalidPassword: "Wrong username or password",
	CodeServerBusy:      "Interval error",
}

func (rescode ResCode) Msg() string {
	msg, ok := codeMsgMap[rescode]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
