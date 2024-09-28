package main

import (
	"fmt"

	"rsc.io/quote"
)

const PI = 3.14

func variableInGo() {
	var name string = "Jiwan"

	age := 20.5

	var (
		home = "Nepal"
		city = "Kathmandu"
	)

	Age := 21

	fmt.Println("Println", name, age, home, city, Age)
	fmt.Printf("My name is %v Chaudhary and age is %v and type of name is %T and age is %T", name, age, name, age)
	fmt.Print("\n", name, "\n", age, home, city, Age)
}

func getSchoolParam(school string) string {
	if school == "" {
		school = "Abc"
	}
	fmt.Println(school)
	return school
}

func dataTypes() {
	// Data Types in Go
	isActive := true // boolean
	var isAdmin bool

	name := "Jiwan"
	var mark float32 = 25.3
	var age int = 5

	const PERMANENT_NUMBER int = 162

	if age > 4 {
		isAdmin = true
	}

	var (
		city    string  = "Ktm"
		area    float32 = 14.4
		code    int     = 123
		houseNo int     = 42
	)
	fmt.Println(isActive, isAdmin, name, mark, age, city, area, code, houseNo, PERMANENT_NUMBER)
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(PI)
	variableInGo()
	school := "Jiwan's School"
	getSchoolParam(school)
	dataTypes()
	fmt.Println(quote.Go())
}
