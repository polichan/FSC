package controller

import (
	"fsc/global"
	"fsc/model"
	"fsc/service"
	"fsc/util"
)

func Login(mobile string, password string, phoneType string)  {
	loginStruct := &model.LoginStruct{
		Mobile: mobile,
		Password: password,
		PhoneType: phoneType,
	}
	err := service.Login(loginStruct)
	if err != nil {
		global.FSC_LOG.Error("登录失败", err)
	}else{
		global.FSC_LOG.Info("登录成功")
		util.ShowUserInfo()
	}
}
