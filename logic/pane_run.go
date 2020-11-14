package logic

import (
	"fsc/controller"
	"fsc/global"
)

func RunTarget()  {
	if global.FSC_CONFIG.Panel.RedDot == 0 ||  global.FSC_CONFIG.Panel.GreenDot == 0{
		global.FSC_LOG.Error("请先在 config.yaml 中配置需要跑几个红点以及绿点")
		return
	}
	global.FSC_LOG.Info("当前配置如下：")
	global.FSC_LOG.Infof("目标跑 %d 个红点", global.FSC_CONFIG.Panel.RedDot)
	global.FSC_LOG.Infof("目标跑 %d 个绿点", global.FSC_CONFIG.Panel.GreenDot)
	controller.RunTarget()
}

func FreeRun()  {

}