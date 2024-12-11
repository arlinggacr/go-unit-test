package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	result := HelloWorld("ara")
	if result != "hello data" {
		// t.FailNow()
		t.Fatal("Result Must be hello data")
	}

	fmt.Println("Test Hello Done")
}

func TestHelloSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		fmt.Println("cannot run in mac os")
	}

	result := HelloWorld("ara")
	require.Equal(t, result, "hi ara")

	fmt.Println("Test Hello Done")
}

func TestHelloAssert(t *testing.T) {
	result := HelloWorld("ara")
	assert.Equal(t, "hello ara", result, "result must be equal") // sama dengan t.Fail()

	fmt.Println("Test Hello Done")
}

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("ara")
	require.Equal(t, "hello ara", result, "result must be equal") // sama dengan t.FailNow()

	fmt.Println("Test Hello Done")
}
