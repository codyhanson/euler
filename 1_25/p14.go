package main

/*
Longest Collatz sequence
Problem 14
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

import (
	"fmt"
)

//returns len of sequence
func collatzLen(start int) int {
	cur := start
	len := 1
	for cur != 1 {
		len++
		if cur%2 == 0 {
			cur = cur / 2
		} else {
			cur = 3*cur + 1
		}

	}
	return len
}

func main() {
	longestLen := 0
	longestStart := 1
	for i := 2; i < 1000000; i++ {
		fmt.Println(i)
		curLen := collatzLen(i)
		if curLen > longestLen {
			longestLen = curLen
			longestStart = i
		}
	}
	fmt.Println(longestStart)
	fmt.Println(longestLen)
}
