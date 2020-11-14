package logic

import (
	"fsc/controller"
	"fsc/global"
	"fsc/util"
)

func Login(item *PanelOptionItem)  {
	if !hasLogin() {
		// 没有登陆
		// 是否在配置文件写了用户信息
		hasInputCorrectAccountInfoInConfig := hasInputCorrectAccountInfoInConfig()
		if !hasInputCorrectAccountInfoInConfig {
			// 如果没填写，告知用户
			global.FSC_LOG.Error("请先在 config.yaml 中设置正确的高校体育账号信息")
			return
		}
		// 有配置信息尝试登陆
		hasLoginSuccess := controller.Login(global.FSC_CONFIG.Account.Mobile, global.FSC_CONFIG.Account.Password, global.FSC_CONFIG.Account.PhoneType)
		if hasLoginSuccess {
			resetItemTitle(item)
			util.ShowUserInfo()
			return
		}
	}else{
		resetItemTitle(item)
		util.ShowUserInfo()
	}
}

func hasInputCorrectAccountInfoInConfig()bool  {
	return global.FSC_CONFIG.Account.Mobile != "" && global.FSC_CONFIG.Account.Password != "" && global.FSC_CONFIG.Account.PhoneType != ""
}

func hasLogin() bool {
	return global.FSC_USER != nil
}

func resetItemTitle(item *PanelOptionItem)  {
	item.Title = "个人资料"
}
