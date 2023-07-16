package errorx

import "google.golang.org/grpc/status"

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return e.Msg
}

var (
	ErrMsgNotSec      = status.Error(10001, "Failed in Wechat MsgSecCheck")
	ErrWechatApiWrong = status.Error(10002, "Call Wechat Api Incorrectly")
	ErrPhotoNotSec    = status.Error(10003, "Fail in Photo Audit ")
)
