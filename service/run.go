package service

import (
	"encoding/json"
	"fsc/global"
	"fsc/global/response"
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
		mp := make(map[string]string)
		_ = util.Transfer(loginRequestStruct, &mp)
		client := util.NewFSCHttpClientSend(util.GetUrlBuild(util.GenerateUrl("/api/run/runPage"), mp))
		body, err := client.Get()
		if err != nil {
			return err
		}else{
			// 红点和绿点
			redDot, greenDot := global.FSC_CONFIG.Panel.RedDot, global.FSC_CONFIG.Panel.GreenDot
			var runMapResponse response.SCRunMapResponseStruct
			err = json.Unmarshal(body, &runMapResponse)
			if err != nil {
				return err
			}
			var (
				possibleBNode []model.IBeaconStruct
				possibleTNode []model.LocationStruct
			)

			for _, item := range runMapResponse.Data.IBeacon{
				bNode := util.Point{Lat: stringToFloat64(item.Position.Latitude), Lon: stringToFloat64(item.Position.Longitude)}
				metre := bNode.MetresTo(util.Point{Lat: loc.Lat, Lon: loc.Lng})
				if metre / 1000 < 60 {
					possibleBNode = append(possibleBNode, item)
				}
			}
			for _, item := range runMapResponse.Data.GPSInfo{
				tNode := util.Point{Lat: stringToFloat64(item.Latitude), Lon: stringToFloat64(item.Longitude)}
				metre := tNode.MetresTo(util.Point{Lat: loc.Lat, Lon: loc.Lng})
				if metre / 1000 < 60 {
					possibleTNode = append(possibleTNode, item)
				}
			}
			var runTargetRequest model.RunStruct
			runTargetRequest.BNode = possibleBNode[:redDot]
			runTargetRequest.TNode = possibleTNode[:greenDot]
			positionInfo := runTargetRequest.BNode[0].Position
			// 起始点
			startPoint := util.NewGPSPoint(stringToFloat64(positionInfo.Latitude), stringToFloat64(positionInfo.Longitude))
			var gpsPointList []*util.GPSPointStruct
			for _, _ = range []int{0, 1} {
				// 起始点进行 Walk，按照 strip 的区间 [-strip, strip] 随机增加或减少步数， 并返回一个新的 point 结构体
				gpsPointList = append(gpsPointList, startPoint.Walk(0.003))
			}
			for _, node := range runTargetRequest.BNode{
				pos := node.Position
				pos.Speed = 0.0
				node.Position = pos
				gpsPointList = append(gpsPointList, util.NewGPSPoint(stringToFloat64(pos.Latitude), stringToFloat64(pos.Longitude)))
			}
			for _, node := range runTargetRequest.TNode{
				node.Speed = 0.0
				gpsPointList = append(gpsPointList, util.NewGPSPoint(stringToFloat64(node.Latitude), stringToFloat64(node.Longitude)))
			}
			// todo 自动根据 GpsPointList 进行寻路
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

func stringToFloat64(v string)float64  {
	float, _ :=strconv.ParseFloat(v, 64)
	return float
}