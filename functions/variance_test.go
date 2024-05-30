package functions

import "testing"

func TestCalculateVariance(t *testing.T) {
	
	tests := []struct {
		name string
		args []float64
		want float64
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateVariance(tt.args); got != tt.want {
				t.Errorf("CalculateVariance() = %v, want %v", got, tt.want)
			}
		})
	}
}
