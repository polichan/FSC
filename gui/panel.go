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
		// 绘制面板
		drawPanel()
		// 根据面板序号执行动作
		execAction(&option)
	}
}

func login()  {
	if isLogin() {
		global.FSC_LOG.Error("您已登录，请勿重复登录")
		return
	}
	accountStruct := global.FSC_CONFIG.Account
	if accountStruct.Mobile == "" || accountStruct.Password == "" || accountStruct.PhoneType == "" {
		global.FSC_LOG.Error("请先在 config.yaml 中设置正确的高校体育账号信息")
	}else{
		controller.Login(accountStruct.Mobile, accountStruct.Password, accountStruct.PhoneType)
	}
}

func isLogin() bool {
	return global.FSC_USER != nil
}

func drawPanel()  {
	fmt.Println("\n                  FSC v2.0")
	fmt.Println("\n***********************************************")
	if isLogin(){
		fmt.Println("\n*          0----------已     登    录          *")
	}else{
		fmt.Println("\n*          0----------账  号   登  录          *")
	}
	fmt.Println("\n*          1----------目    标     跑          *")
	fmt.Println("\n*          2----------自    由     跑          *")
	fmt.Println("\n*          3----------退           出          *")
	fmt.Println("\n***********************************************")
	fmt.Printf("\n 请选择选项号: ")
}

func execAction(option *int)  {
	o := *option
	_, _ = fmt.Scanf("%d", &o)
	switch o {
	case FSCPanelOptionRunTarget:
		controller.RunTarget()
		break
		case FSCPanelOptionFreeRun:
			controller.FreeRun()
		break
		case FSCPanelOptionLogin:
			login()
	}
}