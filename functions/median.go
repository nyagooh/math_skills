package functions

import (
	"sort"
)

// calculate the median of a data set
func CalculateMedian(data []float64) float64 {

	if len(data) == 0 {
		return 0.0
	}
	sort.Float64s(data)
	output := 0.0
	result := len(data) / 2
	if len(data)%2 == 0 {
		output = (data[result] + data[result-1]) / 2.0
	} else {
		output = data[result]
	}
	return output
}
