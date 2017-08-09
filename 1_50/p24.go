package main

/*Lexicographic permutations
Problem 24
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

import (
	"bytes"
	"fmt"
)

func concat(digits []string) string {
	var currentP bytes.Buffer
	for j := 0; j < len(digits); j++ {
		currentP.WriteString(digits[j])
	}
	return currentP.String()
}

var permutations = make([]string, 1000000, 1000000)

//https://en.wikipedia.org/wiki/Heap%27s_algorithm
func permute(n int, s []string) {
	if n == 1 {
		permutations = permutations.append(s)
	}
}

func main() {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	count := 0

	permute(10, digits)

	fmt.Println("-------")
	fmt.Println(permutations[1000000])
}
