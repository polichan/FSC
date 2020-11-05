package service

import (
	"encoding/json"
	"fsc/model"
	"fsc/model/request"
	"fsc/util"
	uuid "github.com/satori/go.uuid"
)


func Login(loginModel *model.LoginStruct)(err error)  {
	// 添加 uuid
	loginModel.Info = uuid.NewV4()
	// 序列化
	v, _ := json.Marshal(loginModel)
	// 获取 sign
	sign := util.GetMd5(v)
	// 请求体
	var loginRequestStruct request.LoginRequestStruct
	loginRequestStruct.Sign = sign
	loginRequestStruct.Data = string(v)
	// 转 map
	mp:=make(map[string]string)
	_ = util.Transfer(loginRequestStruct, &mp)
	// 请求
	client := util.NewFSCHttpClientSend(util.GetUrlBuild(util.GenerateUrl("/api/reg/login"), mp))
	_, err = client.Get()
	return err
}