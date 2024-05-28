package functions


func CalculateMean(data []float64)float64 {
	// sort.float64(data)
	number := 0.0
	output := 0.0
	for _, ch := range data {
		number += ch
	}
	output = number / float64(len(data))
	return output
}