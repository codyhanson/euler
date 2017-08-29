package main

/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

import (
	"fmt"
)

const (
	SUN int = iota
	MON
	TUE
	WED
	THU
	FRI
	SAT
)

var monthDayMap = map[int]int{
	0:  31,
	1:  28,
	2:  31,
	3:  30,
	4:  31,
	5:  30,
	6:  31,
	7:  31,
	8:  30,
	9:  31,
	10: 30,
	11: 31,
}

func nextDay(day int) int {
	return (day + 1) % 7
}

func nextMonth(month int) int {
	return (month + 1) % 12
}

func next(day, dom, m, y int) (oDay, oDom, oM, oY int) {
	oDay = nextDay(day)
	oM = m //assume no change
	oY = y

	if dom == monthDayMap[m] && m != 1 {
		//start new month
		oM = nextMonth(m)
		oDom = 1
	} else if m == 1 && dom == monthDayMap[1] {
		if y%4 == 0 {
			//leapyear
			oDom = dom + 1
		} else {
			oDom = 1
			oM = nextMonth(m)
		}
	} else if m == 1 && dom == 29 {
		//leap day over now.
		oDom = 1
		oM = nextMonth(m)
	} else {
		oDom = dom + 1
	}

	if m == 11 && dom == monthDayMap[11] {
		//end of year.
		oY = y + 1
	}

	return
}

func main() {

	count := 0
	day, dom, m, y := MON, 1, 0, 1900
	for {
		fmt.Printf("%d, %d, %d, %d\n", day, dom, m, y)
		day, dom, m, y = next(day, dom, m, y)
		if day == SUN && dom == 1 {
			count++
		}
		if dom == 31 && y == 2000 {
			break
		}
	}

	fmt.Println("=========")
	fmt.Println(count)
}
