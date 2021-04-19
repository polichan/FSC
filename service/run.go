package service

import (
	"encoding/json"
	"fmt"
	"fsc/global"
	"fsc/global/response"
	"fsc/model"
	"fsc/model/request"
	"fsc/util"
	"time"
)

// 体育锻炼
func RunTarget(user *model.UserStruct) (err error) {
	loc, err := util.GetSchoolLocation(user.School)
	if err != nil {
		return err
	} else {
		// params
		initLocation := util.Float64ToString(loc.Lng) + "," + util.Float64ToString(loc.Lat)
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
		client := util.NewFSCHttpClientSend(util.GetUrlBuilder(util.GenerateUrl("/api/run/runPage"), mp))
		body, err := client.Get()
		if err != nil {
			return err
		} else {
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
			for _, item := range runMapResponse.Data.IBeacon {
				bNode := util.Point{Lat: util.StringToFloat64(item.Position.Latitude), Lon: util.StringToFloat64(item.Position.Longitude)}
				metre := bNode.MetresTo(util.Point{Lat: loc.Lat, Lon: loc.Lng})
				if metre/1000 < 60 {
					possibleBNode = append(possibleBNode, item)
				}
			}
			for _, item := range runMapResponse.Data.GPSInfo {
				tNode := util.Point{Lat: util.StringToFloat64(item.Latitude), Lon: util.StringToFloat64(item.Longitude)}
				metre := tNode.MetresTo(util.Point{Lat: loc.Lat, Lon: loc.Lng})
				if metre/1000 < 60 {
					possibleTNode = append(possibleTNode, item)
				}
			}
			var runTargetRequest model.RunStruct
			runTargetRequest.BNode = possibleBNode[:redDot]
			runTargetRequest.TNode = possibleTNode[:greenDot]
			positionInfo := runTargetRequest.BNode[0].Position
			// 起始点
			startPoint := util.NewGPSPoint(util.StringToFloat64(positionInfo.Latitude), util.StringToFloat64(positionInfo.Longitude))
			var gpsPointList []*util.GPSPointStruct
			for range []int{0, 1} {
				// 起始点进行 Walk，按照 strip 的区间 [-strip, strip] 随机增加或减少步数， 并返回一个新的 point 结构体
				gpsPointList = append(gpsPointList, startPoint.Walk(0.003))
			}
			for _, node := range runTargetRequest.BNode {
				pos := node.Position
				pos.Speed = 0.0
				node.Position = pos
				gpsPointList = append(gpsPointList, util.NewGPSPoint(util.StringToFloat64(pos.Latitude), util.StringToFloat64(pos.Longitude)))
			}
			for _, node := range runTargetRequest.TNode {
				node.Speed = 0.0
				gpsPointList = append(gpsPointList, util.NewGPSPoint(util.StringToFloat64(node.Latitude), util.StringToFloat64(node.Longitude)))
			}
			//path := util.PathPlan(gpsPointList)
			// todo 自动根据 GpsPointList 进行寻路
			distance := 2.0
			speed := util.RandomInt(300, 500) // 每千米需要花费多少秒钟
			duration := int(distance) * speed
			runTargetRequest.StartTime = time.Now().Format("2006-01-02 15:04:05")
			runTargetRequest.EndTime = time.Now().Add(time.Duration(duration)).Format("2006-01-02 15:04:05")
			runTargetRequest.UserId = util.StringToInt(user.UserId)
			runTargetRequest.RunPageId = util.IntToString(runMapResponse.Data.RunPageId)
			runTargetRequest.Real = util.Float64ToString(distance * 1000)
			runTargetRequest.Duration = util.IntToString(duration)
			runTargetRequest.Speed = util.IntToString(speed)
			// todo 寻路生成的 track runTargetRequest.Track
			// 配速 =  1000 / (步频 * 步幅)
			runTargetRequest.BuPin = fmt.Sprintf("%.1f", util.RandomFloat(120.0, 140.0))
			runTargetRequest.TotalNum = util.RandomInt(2000, 3000)
			global.FSC_LOG.Infof("计划跑步 %f km，本工具将在 %s 自动提交跑步记录，未完成前，请勿退出。", distance, runTargetRequest.EndTime)
			time.Sleep(time.Duration(duration))
			err = saveRun(runTargetRequest)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// 自由跑
func FreeRun(user *model.UserStruct) (err error) {
	return nil
}

func saveRun(data model.RunStruct) (err error) {
	v, _ := json.Marshal(data)
	sign := util.GetMd5(v)
	var saveRunRequest request.SCRequestStruct
	saveRunRequest.Sign = sign
	saveRunRequest.Data = string(v)
	mp := make(map[string]string)
	_ = util.Transfer(saveRunRequest, &mp)
	client := util.NewFSCHttpClientSend(util.GetUrlBuilder(util.GenerateUrl("/api/run/saveRunV2"), mp))
	_, err = client.Get()
	if err != nil {
		return err
	}
	return nil
}
