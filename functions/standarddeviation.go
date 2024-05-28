package functions

import (
	"math"
)

func CalculateStandardSDeviation(data []float64)float64{
	number := 0.0
	for _, ch := range data {
		number += ch - CalculateMean(data) * ch - CalculateMean(data)
	}
	initial := number/ float64(len(data))
	return math.Sqrt(initial)
	// for i := 1;i < 1000;i++ {
	// initial = initial - (initial*initial-number)/(2*initial)
	// }
	// return initial
}