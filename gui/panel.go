package gui

import (
	"fmt"
	"fsc/controller"
	"fsc/global"
)

const (
	FSCPanelOptionLogin = iota
	FSCPanelOptionRunTarget
	FSCPanelOptionFreeRun
	FSCPanelOptionExit
)

func ShowPanel()  {
	option := 0
	for option != FSCPanelOptionExit {
		fmt.Println("\n\t\t                  FSC v2.0")
        fmt.Println("\n\t\t***********************************************")
        fmt.Println("\n\t\t*          0----------高 校 账 号 登 录         *")
        fmt.Println("\n\t\t*          1----------目    标     跑          *")
        fmt.Println("\n\t\t*          2----------自    由     跑          *")
        fmt.Println("\n\t\t*          3----------退           出          *")
        fmt.Println("\n\t\t***********************************************")
		fmt.Printf("\n\n\t\t 请选择菜单号:  ")
		_, _ = fmt.Scanf("%d", &option)
		switch option {
			case FSCPanelOptionRunTarget:
			break
			case FSCPanelOptionFreeRun:
			break
			case FSCPanelOptionLogin:
			Login()
		}
	}
}

func Login()  {
	accountStruct := global.FSC_CONFIG.Account
	if accountStruct.Mobile == "" || accountStruct.Password == "" || accountStruct.PhoneType == "" {
		global.FSC_LOG.Error("请先在 config.yaml 中设置正确的高校体育账号信息")
	}else{
		controller.Login(accountStruct.Mobile, accountStruct.Password, accountStruct.PhoneType)
	}
}
