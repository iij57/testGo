package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(hello())
		time.Sleep(10 * time.Second)
	}

}

func hello() string {
	return "Hello!!!"

}
