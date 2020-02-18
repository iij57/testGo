package main

import (
	"main"
	"testing"
)

func TestHello(t *testing.T) {
	h := main.Hello()

	if h != "hello!!!" {
		t.Error("Wrong Result")
	}
}
