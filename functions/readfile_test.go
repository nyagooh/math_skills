package functions

import (
	"reflect"
	"testing"
)

func TestReadfile(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []float64
	}{
		{"dataset", "data", []float64{189,113,121,114,145,122}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Readfile(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Readfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
