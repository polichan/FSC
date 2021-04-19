package util

import "strconv"

func Float64ToString(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 32)
}

func StringToFloat64(v string) float64 {
	float, _ := strconv.ParseFloat(v, 64)
	return float
}

func IntToString(v int) string {
	return strconv.Itoa(v)
}

func StringToInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}
