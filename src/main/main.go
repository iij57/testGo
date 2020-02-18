package main

import (
	"calc"
	"fmt"
)

func main() {
	fmt.Println(Hello())
	fmt.Println(calc.Sum(2, 3))
}

func Hello() string {
	return "Hello!!!"

}
