package helper

import "testing"

func TestCalculateData(t *testing.T) {
	result := CalculateData(1, 1)
	if result != 2 {
		panic("result must be 2")
	}
}
