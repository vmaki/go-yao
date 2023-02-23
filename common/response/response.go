package response

import (
	"encoding/json"
)

type RespData struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func New(code ResCode) *RespData {
	return &RespData{
		Code: code,
		Msg:  code.Msg(),
	}
}

func (e *RespData) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}
