package functions

import "sort"

func CalculateMedian(data []float64) float64 {
	sort.Float64s(data)
	output := 0.0
	result := len(data) / 2
	if len(data)%2 == 0 {
		output = (data[result] + data[result-1]) / 2
		return output
	}
	output = data[result]
	return output
}
