package main

import (
	"fmt"
	"time"
)

var name string = "Jiwan"

var newName string = "Jiwan new"

func main() {

	age := 18

	if age > 20 {
		fmt.Println("not eligible")
	} else if age < 17 {
		fmt.Println("Eligible")
	} else {
		fmt.Println("what the f")
	}
	fmt.Println(time.Date(2024, time.Month(4), 23, 12, 53, 23, 123, time.Local), "current")

	fmt.Println("new print")
	fmt.Println(name, newName)
}
