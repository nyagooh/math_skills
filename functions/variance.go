package functions
//calculate variance of a data set
func CalculateVariance(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}
	var number float64
	for _, ch := range data {
		number += (ch - CalculateMean(data))*(ch - CalculateMean(data))
	}
	return number / float64(len(data))
}
