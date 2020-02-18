package test

import (
	"calc"
	"testing"
)

func TestHello(t *testing.T) {
	result := calc.Sum(2, 3)

	if result != 5 {
		t.Error("Wrong Result")
	}
}
