package response

import (
	"encoding/json"
)

type RespData struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func New(code ResCode, msg ...string) *RespData {
	return &RespData{
		Code: code,
		Msg:  defaultMessage(code.Msg(), msg...),
	}
}

func (e *RespData) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}
