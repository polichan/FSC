package util

import (
	"math/rand"
)

type IGPSPoint interface {
	Walk()
}

type GPSPointStruct struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

func NewGPSPoint(lat, lng float64)*GPSPointStruct  {
	return &GPSPointStruct{
		Latitude: lat,
		Longitude: lng,
	}
}

func (G *GPSPointStruct) Walk(strip float64)*GPSPointStruct {
	return NewGPSPoint(G.Latitude + random(-strip, strip), G.Longitude + random(-strip, strip))
}


func random(min, max float64)float64  {
	return min + rand.Float64() * (max - min)
}


