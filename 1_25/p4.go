package main

/*
Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

import "fmt"

func numberIsPalindrome(i int) bool {
	//TODO
	return false
}

func main() {

	largestPalindrome := 0

	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			p := i * j
			if numberIsPalindrome(p) && p > largestPalindrome {
				largestPalindrome = p
			}
		}
	}
	fmt.Println(largestPalindrome)
}
