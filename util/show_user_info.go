package util

import "fsc/global"

func ShowUserInfo()  {
	global.FSC_LOG.Info("当前登录账户信息：")
	global.FSC_LOG.Info("用户名：", global.FSC_USER.Username)
	global.FSC_LOG.Info("所在学校：", global.FSC_USER.School)
	global.FSC_LOG.Info("所属课程：", global.FSC_USER.Course)
	global.FSC_LOG.Info("课程教师：", global.FSC_USER.Teacher)
	global.FSC_LOG.Info("目标公里数：", global.FSC_USER.Goal, " KM")
	global.FSC_LOG.Info("剩余公里数：", global.FSC_USER.Surplus, " KM")
	global.FSC_LOG.Info("用户 Token：", global.FSC_USER.UToken)
}