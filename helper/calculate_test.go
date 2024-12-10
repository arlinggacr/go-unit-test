package helper

import (
	"fmt"
	"testing"
)

func TestCalculateData(t *testing.T) {
	result := CalculateData(1, 1)
	if result != 3 {
		// t.Fail()
		t.Error("Calculate data must be 3")
	}

	fmt.Println("test CalculateData Done")
}
