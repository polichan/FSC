package gui

import (
	"fsc/logic"
	"os"
)

func ShowPanel()  {
	// 构建面板
	panel := logic.NewFSCPanel()
	// 设置面板选项数据
	panel.SetOptions(PrepareOptionItems())
	// 绘制面板
	panel.Draw()
	// 询问用户进行选择
	panel.Ask()
	// 监听用户选择
	panel.WatchingOption()
}

func PrepareOptionItems()[]logic.PanelOptionItem  {
	// 创建数据源
	panelOptionOne := logic.NewPanelOptionItem(logic.FSCPanelOptionLogin, func() {
		logic.Login()
	}, "账号登录")
	panelOptionTwo := logic.NewPanelOptionItem(logic.FSCPanelOptionRunTarget, func() {
		logic.RunTarget()
	}, "体育锻炼")
	panelOptionThree := logic.NewPanelOptionItem(logic.FSCPanelOptionFreeRun, func() {
		logic.FreeRun()
	}, "自由跑步")
	panelOptionFour := logic.NewPanelOptionItem(logic.FSCPanelOptionExit, func() {
		os.Exit(0)
	}, "退出系统")
	return []logic.PanelOptionItem{
		panelOptionOne,
		panelOptionTwo,
		panelOptionThree,
		panelOptionFour,
	}
}
