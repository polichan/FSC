package controller

import (
	"fsc/global"
	"fsc/service"
)

func RunTarget() {
	err := service.RunTarget(global.FSC_USER)
	if err != nil {
		global.FSC_LOG.Error("体育锻炼失败：", err)
		return
	}
	global.FSC_LOG.Info("体育锻炼成功")
}

func FreeRun() {
	err := service.FreeRun(global.FSC_USER)
	if err != nil {
		global.FSC_LOG.Error("自由跑失败：", err)
		return
	}
	global.FSC_LOG.Info("自由跑成功")
}
