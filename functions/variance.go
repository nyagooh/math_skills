package functions


func CalculateVariance(data []float64)float64{
	number := 0.0
	for _, ch := range data {
		number += (ch - CalculateMean(data) * ch - CalculateMean(data))
	}
  return number / float64(len(data))
}