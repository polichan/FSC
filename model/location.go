package model

// 学校地址
type SchoolLocationStruct struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type LocationStruct struct {
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
}