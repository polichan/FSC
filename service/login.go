package service

import (
	"encoding/json"
	"fsc/global"
	"fsc/global/response"
	"fsc/model"
	"fsc/model/request"
	"fsc/util"
	"github.com/mitchellh/mapstructure"
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
	var loginRequestStruct request.SCRequestStruct
	loginRequestStruct.Sign = sign
	loginRequestStruct.Data = string(v)
	// 转 map
	mp:=make(map[string]string)
	_ = util.Transfer(loginRequestStruct, &mp)
	// 请求
	client := util.NewFSCHttpClientSend(util.GetUrlBuild(util.GenerateUrl("/api/reg/login"), mp))
	body, err := client.Get()
	if err != nil {
		return err
	}
	scRep := response.SCResponse{}
	scUser := model.UserStruct{}
    err = json.Unmarshal(body, &scRep)
    // 传递给全局
    err = mapstructure.Decode(scRep.Data, &scUser)
    global.FSC_USER = &scUser
	return err
}