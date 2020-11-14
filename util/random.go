package util

import "math/rand"

func RandomFloat(min, max float64)float64  {
	return min + rand.Float64() * (max - min)
}

func RandomInt(min, max int)int  {
	return rand.Intn(max-min) + min
}

