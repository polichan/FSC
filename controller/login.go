package controller

import (
	"fsc/global/response"
	"fsc/model"
	"fsc/service"
)

func Login(mobile string, password string, phoneType string)  {
	loginStruct := &model.LoginStruct{
		Mobile: mobile,
		Password: password,
		PhoneType: phoneType,
	}
	err := service.Login(loginStruct)
	if err != nil {
		response.FailWithMessage("登录失败", err)
	}else{
		response.OkWithMessage("登录成功")
	}
}