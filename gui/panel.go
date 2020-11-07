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


func PrepareOptionItems()[]*logic.PanelOptionItem  {
	var (
		panelOptionOne,
		panelOptionTwo,
		panelOptionThree,
		panelOptionFour logic.PanelOptionItem
	)
	// 创建数据源
	panelOptionOne.SetPanelOptionItem(logic.FSCPanelOptionLogin, func() {
		logic.Login(&panelOptionOne)
	}, "账号登录")

	panelOptionTwo.SetPanelOptionItem(logic.FSCPanelOptionRunTarget, func() {
		logic.RunTarget()
	}, "体育锻炼")

	panelOptionThree.SetPanelOptionItem(logic.FSCPanelOptionFreeRun, func() {
		logic.FreeRun()
	}, "自由跑步")

	panelOptionFour.SetPanelOptionItem(logic.FSCPanelOptionExit, func() {
		os.Exit(0)
	}, "退出系统")

	return []*logic.PanelOptionItem{
		&panelOptionOne,
		&panelOptionTwo,
		&panelOptionThree,
		&panelOptionFour,
	}
}
