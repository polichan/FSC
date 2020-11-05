package model

// 打卡跑
type RunStruct struct {
	BNode []int `json:"bNode"`
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
	TNode []int `json:"tNode"`
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