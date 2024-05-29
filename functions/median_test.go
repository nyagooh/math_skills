package functions

import "testing"

func TestCalculateMedian(t *testing.T) {

	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{"odd numbers", []float64{189, 113, 121, 114, 145}, 121},
		{"even numbers", []float64{189, 113, 121, 114, 145,122}, 122},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMedian(tt.args); got != tt.want {
				t.Errorf("CalculateMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
