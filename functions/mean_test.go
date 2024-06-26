package functions

import "testing"

func TestCalculateMean(t *testing.T) {
	
	tests := []struct {
		name string
		input []float64
		want float64
	}{
		{"dataset", []float64{189,113,121,114,145,122},134},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMean(tt.input); got != tt.want {
				t.Errorf("CalculateMean() = %v, want %v", got, tt.want)
			}
		})
	}
}
