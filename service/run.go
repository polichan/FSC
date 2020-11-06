package service

import (
	"encoding/json"
	"fmt"
	"fsc/model"
	"fsc/model/request"
	"fsc/util"
	"strconv"
)

// 体育锻炼
func RunTarget(user *model.UserStruct)(err error)  {
	loc, err := util.GetSchoolLocation(user.School)
	if err != nil {
		return err
	}else{
		// params
		initLocation := float64ToString(loc.Lng) + "," + float64ToString(loc.Lat)
		// 体育锻炼参数
		data := make(map[string]string, 3)
		data["initLocation"] = initLocation
		data["type"] = "1"
		data["userid"] = user.UserId
		v, _ := json.Marshal(data)
		// 签名
		sign := util.GetMd5(v)
		var loginRequestStruct request.SCRequestStruct
		loginRequestStruct.Sign = sign
		loginRequestStruct.Data = string(v)
		mp:=make(map[string]string)
		_ = util.Transfer(loginRequestStruct, &mp)
		client := util.NewFSCHttpClientSend(util.GetUrlBuild(util.GenerateUrl("/api/run/runPage"), mp))
		body, err := client.Get()
		if err != nil {
			return err
		}else{
			// todo 分析结果坐标，自动寻路
			fmt.Printf("%s", string(body))
		}
	}
	return err
}

// 自由跑
func FreeRun(user *model.UserStruct)(err error)  {
	return nil
}

func float64ToString(v float64)string  {
	return strconv.FormatFloat(v, 'f', -1, 32)
}