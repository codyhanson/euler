package main

import (
	"eulertools"
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if eulertools.IsPrime(i) {
			fmt.Println(i)
		}
	}
}
