package main

/*
Non-abundant sums
Problem 23
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

import (
	"fmt"
)

func isAbundant(n int) bool {
	sum := 0

	//find and sum divisors
	for i := 1; i < n; i++ {
		if n%i == 0 {
			//divisor
			sum += i
		}
	}
	return sum > n
}

func canSumFromAbundant(n int) bool {
	//reaaaaally innefficient
	for i := 0; i < len(abundants); i++ {
		for j := i; j < len(abundants); j++ {
			if abundants[i]+abundants[j] == n {
				fmt.Printf("sum: %d + %d = %d\n", abundants[i], abundants[j], n)
				return true
			}
		}
	}
	return false
}

var abundants []int

func main() {
	sum := 0
	for i := 1; i <= 28123; i++ {
		//first find abundants in this range
		if isAbundant(i) {
			abundants = append(abundants, i)
		}
	}

	for n := 1; n <= 28123; n++ {
		if !canSumFromAbundant(n) {
			sum += n
		}
	}

	fmt.Println(sum)
}
