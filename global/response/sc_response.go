package response

import "fsc/model"

type SCResponse struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

type SCRunMapResponseStruct struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
	Data model.RunMapStruct `json:"data"`
}