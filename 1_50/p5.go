package main

/*
Smallest Multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

import "fmt"

func main() {
	n := 20
	for {
		success := true
		for i := 2; i < 21; i++ {
			if n%i != 0 {
				success = false
				break
			}
		}

		if success {
			fmt.Println(n)
			return
		}
		n += 1
	}

}
