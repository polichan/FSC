package service

import (
	"fsc/model"
	"fsc/util"
)

// 体育锻炼
func RunTarget(user *model.UserStruct)(err error)  {
	err = util.GetSchoolLocation("上海电机学院")
	return err
}

// 自由跑
func FreeRun(user *model.UserStruct)(err error)  {
	return nil
}