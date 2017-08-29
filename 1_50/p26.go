package main

import (
	"fmt"
	"strconv"
)

/*
A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/

func cycleLen(d float64) int {
	s := strconv.FormatFloat(d, 'f', 100, 64)
	fmt.Println(s)
	return 0
}

func main() {
	var d float64
	longest := 0
	var longestD float64
	for d = 2; d < 1000; d++ {
		f := 1 / d
		l := cycleLen(f)
		if l > longest {
			longest = l
			longestD = d
		}
		fmt.Printf("%v: %v\n", d, f)
	}
	fmt.Println("=============")
	fmt.Printf("%v: %v %v\n", longestD, longest, 1/longestD)
}
