package utils

import "math"

func RoundDown(val float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Floor(val*factor) / factor
}
