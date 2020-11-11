package util

import (
	"math"
)

var earthRadiusMetres float64 = 6371000

type Metres float64

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

type Point struct {
	Lat float64
	Lon float64
}

func (p Point) MetresTo(remote Point) Metres {
	return Distance(p, remote)
}

func (p Point) Delta(point Point) Delta {
	return Delta{
		Lat: p.Lat - point.Lat,
		Lon: p.Lon - point.Lon,
	}
}

func (p Point) toRadians() Point {
	return Point{
		Lat: degreesToRadians(p.Lat),
		Lon: degreesToRadians(p.Lon),
	}
}

type Delta struct {
	Lat float64
	Lon float64
}

func Distance(origin, position Point) Metres {
	origin = origin.toRadians()
	position = position.toRadians()

	change := origin.Delta(position)

	a := math.Pow(math.Sin(change.Lat/2), 2) + math.Cos(origin.Lat)*math.Cos(position.Lat)*math.Pow(math.Sin(change.Lon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return Metres(earthRadiusMetres * c)
}
