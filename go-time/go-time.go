package main

import (
	"fmt"
	"time"
)

type Days int

const (
	Sunday Days = iota
	Monday
	Tuesday
	Wednesday
	Thrusday
	Friday
	Saturday
)

type Color int

const (
	Red Color = 19 << iota
	Green
	Yellow
	Orange
	Blue
)

func yourColor(c Color) string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Orange:
		return "Orange"
	case Yellow:
		return "Yellow"
	default:
		return "Null"
	}
}

func isWeekend() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("Week day")
	}
}

func isEligible(age int) {

	if age > 18 {
		fmt.Println("eligible")
	} else {
		fmt.Println("Not eligible")
	}

	switch age {
	case age:
		fmt.Println("Not eligible")
	default:
		fmt.Println("eligible")
	}
}

func newFunc(a, b int) int {
	return a + b
}

func WhatDay(d int) string {
	switch d {
	case 1:
		return "One"
	default:
		return "WTF"
	}
}

func main() {
	isWeekend()
	// result := newFunc(2, 4)
	println(newFunc(2, 4))
	isEligible(52)

	var myColor Color = Color(Blue)
	result := yourColor(myColor)

	fmt.Println(result, "color")
	fmt.Println(Days(Friday), "days")

	var day Days = Days(Blue)
	fmt.Println(day, "1228----")
	results := WhatDay(int(day))
	fmt.Println(results, "sh---")

}
