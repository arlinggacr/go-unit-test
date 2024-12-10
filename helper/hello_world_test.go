package helper

import "testing"

func TestHello(t *testing.T) {
	result := HelloWorld("ara")
	if result != "hello ara" {
		panic("result must be hello ara")
	}
}
