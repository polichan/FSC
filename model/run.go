package model

import (
	uuid "github.com/satori/go.uuid"
)

// 打卡跑
type RunStruct struct {
	BNode []IBeaconStruct `json:"bNode"`
	BuPin int `json:"buPin"`
	Busu int `json:"bushu"`
	Duration int `json:"duration"`
	EndTime string `json:"endTime"`
	Frombp int `json:"frombp"`
	Goal []int `json:"goal"`
	Real string `json:"real"`
	RunPageId string `json:"runPageId"`
	Speed string `json:"speed"`
	StartTime string `json:"startTime"`
	TNode []LocationStruct `json:"tNode"`
	TotalNum int `json:"totalNum"`
	Track []int `json:"track"`
	Trend []int `json:"trend"`
	Type string `json:"type"`
	UserId string `json:"userId"`
}

// 自由跑
type FreeRunStruct struct {
	BNode []int `json:"bNode"`
	BuPin int `json:"buPin"`
	Busu int `json:"bushu"`
	Duration int `json:"duration"`
	EndTime string `json:"endTime"`
	Frombp int `json:"frombp"`
	Real string `json:"real"`
	RunPageId string `json:"runPageId"`
	Speed string `json:"speed"`
	StartTime string `json:"startTime"`
	TotalNum int `json:"totalNum"`
	Track []int `json:"track"`
	Trend []int `json:"trend"`
	Type string `json:"type"`
	UserId string `json:"userId"`
}

type RunMapStruct struct {
	BeaconCount int `json:"beaconcount"`
	Gpnt int `json:"gpnt"`
	Distance int `json:"distance"`
	Peisu int `json:"peisu"`
	GPSInfo []LocationStruct `json:"gpsinfo"`
	Length string `json:"length"`
	DayTarget string `json:"dayTarget"`
	IBeacon []IBeaconStruct `json:"ibeacon"`
	RunPageId int `json:"runPageId"`
	MaxSeconds int `json:"maxSeconds"`
}

type IBeaconStruct struct {
	Name string `json:"name"`
	Type int `json:"type"`
	Number string `json:"number"`
	Position LocationStruct `json:"position"`
	ID string `json:"id"`
	UUID uuid.UUID `json:"uuid"`
	Major string `json:"major"`
	Minor string `json:"minor"`
}
