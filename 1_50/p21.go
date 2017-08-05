package main

/*
Amicable numbers
Problem 21
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

import (
	"fmt"
)

var amicables = make(map[int]bool)
var dCache = make(map[int]int)

func d(n int) int {
	s, ok := dCache[n]
	if ok {
		return s
	}
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	dCache[n] = sum
	return sum
}

func amicable(a, b int) bool {
	return d(a) == b && a == d(b)
}

func main() {

	for i := 1; i <= 10000; i++ {
		for j := 1; j <= 10000; j++ {
			if i == j {
				continue
			}
			if amicables[i] && amicables[j] {
				continue
			}
			if amicable(i, j) {
				fmt.Printf("%d %d Amicable\n", i, j)
				amicables[i] = true
				amicables[j] = true
			}
		}
	}

	sum := 0
	for am := range amicables {
		sum += am
	}
	fmt.Println(sum)
}
