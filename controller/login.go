package controller

import (
	"fsc/global"
	"fsc/model"
	"fsc/service"
)

func Login(mobile string, password string, phoneType string) bool {
	loginStruct := &model.LoginStruct{
		Mobile: mobile,
		Password: password,
		PhoneType: phoneType,
	}
	err := service.Login(loginStruct)
	if err != nil {
		global.FSC_LOG.Error("登录失败", err)
		return false
	}else{
		global.FSC_LOG.Info("登录成功")
		return true
	}
}
