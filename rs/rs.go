//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/28 5:32 下午
//
//============================================================

package rs

import (
	"fmt"
)

const (
	Success    = 0
	SuccessMsg = "调用成功"
)

type ApiErr struct {
	ErrCode    int    `json:"errCode"`
	ErrMessage string `json:"errMessage"`
}

func NewApiErr(code int, message string) *ApiErr {
	return &ApiErr{
		ErrCode:    code,
		ErrMessage: message,
	}
}

func (e ApiErr) Error() string {
	return fmt.Sprintf("API异常【%d】%s", e.ErrCode, e.ErrMessage)
}

type Result struct {
	ErrCode    int         `json:"errCode"`
	ErrMessage string      `json:"errMessage"`
	Data       interface{} `json:"data"`
}

func NewResult(data interface{}) Result {
	return Result{
		ErrCode:    Success,
		ErrMessage: SuccessMsg,
		Data:       data,
	}
}

func NewNoDataResult() Result {
	return Result{
		ErrCode:    Success,
		ErrMessage: SuccessMsg,
		Data:       nil,
	}
}
