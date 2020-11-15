package util

import (
	"errors"
	"fsc/global"
	"fsc/model"
	"fsc/sdk/bMap"
)

func GetSchoolLocation(school string) (loc model.SchoolLocationStruct, err error) {
	var schoolLocation model.SchoolLocationStruct
	// 百度 AK 判空
	if global.FSC_CONFIG.BaiduMapConfig.AccessKey == "" {
		return schoolLocation, errors.New("请先配置百度 AccessKey ")
	}
	// 百度 SDK
	bc := bMap.NewBaiduMapClient(global.FSC_CONFIG.BaiduMapConfig.AccessKey)
	// 地址转经纬
	addressToGEO, err := bc.GetGeoViaAddress(school)
	if err != nil {
		return schoolLocation, errors.New("获取学校地址失败")
	}
	if global.FSC_CONFIG.FSC.Debug {
		global.FSC_LOG.Info("学校地址 - Lat", addressToGEO.Result.Location.Lat)
		global.FSC_LOG.Info("学校地址 - Lng", addressToGEO.Result.Location.Lng)
	}
	schoolLocation.Lat = addressToGEO.Result.Location.Lat
	schoolLocation.Lng = addressToGEO.Result.Location.Lng
	return schoolLocation, nil
}


//
//// 路劲规划
//func PathPlan(gpsPointList []*GPSPointStruct)  {
//	index := 0
//	//paths := []int{}
//	//dis := 0
//	for index < len(gpsPointList) - 1 {
//		startPoint, endPoint := gpsPointList[index], gpsPointList[index + 1]
//		// route
//		index +=1
//	}
//	//paths := GenerateHumanLikeRoute(paths)
//}
//
//func GetRoute(){
//	bc := gobaidumap.NewBaiduMapClient(global.FSC_CONFIG.BaiduMapConfig.AccessKey)
//}

//func GenerateHumanLikeRoute()  {
//
//}