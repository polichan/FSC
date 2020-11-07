package logic

import (
	"fsc/controller"
	"fsc/global"
	"fsc/util"
)

func Login()  {
	if IsLogin() {
		util.ShowUserInfo()
		return
	}
	accountStruct := global.FSC_CONFIG.Account
	if accountStruct.Mobile == "" || accountStruct.Password == "" || accountStruct.PhoneType == "" {
		global.FSC_LOG.Error("请先在 config.yaml 中设置正确的高校体育账号信息")
	}else{
		controller.Login(accountStruct.Mobile, accountStruct.Password, accountStruct.PhoneType)
	}
}

func IsLogin() bool {
	return global.FSC_USER != nil
}