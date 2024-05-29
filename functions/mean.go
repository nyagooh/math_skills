package functions
//calculate mean of a data set
func CalculateMean(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}

	number := 0.0
	for _, ch := range data {
		number += ch
	}
	output := number / float64(len(data))
	return output
}
