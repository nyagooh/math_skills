package functions

import "testing"

func TestCalculateMean(t *testing.T) {
	
	tests := []struct {
		name string
		input []float64
		want float64
	}{
		{"Negative numbers", []float64{-1, -2, -3, -4, -5}, -3},
		{"possitive numbers", []float64{189,113,121,114,145,110},132},
		{"decimal numbers", []float64{189.8,113.9,121.6,114.45,145.56,110.34},133},
		
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMean(tt.input); got != tt.want {
				t.Errorf("CalculateMean() = %v, want %v", got, tt.want)
			}
		})
	}
}
