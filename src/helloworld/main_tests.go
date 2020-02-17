package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := hello()
	if result != "Hello!!!" {
		t.Fatalf("hello() didn't output")
	}
}
