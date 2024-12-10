package helper

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	result := HelloWorld("ara")
	if result != "hello data" {
		// t.FailNow()
		t.Fatal("Result Must be hello data")
	}

	fmt.Println("Test Hello Done")
}
