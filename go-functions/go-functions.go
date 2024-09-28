// different ways to define and use functions in Go:
// A function is declared using the func keyword.
// You define the function name, parameters, and return type(s).
package main

import (
	"fmt"
)

// 1. Basic Function Declaration
func add(a int, b int) int {
	c := a + b
	return c
}

// 2. Function with Multiple Return Values
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by 0")
	}
	return a / b, nil
}

// 3. Named Return Values
func swap(a int, b int) (x int) {
	x = b + a
	return
}

// 4. Variadic Functions
func sum(numbers ...int) int {

	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// 7. Functions as First-Class Citizens => Functions can be assigned to variables, passed as arguments, or returned from other functions.

func firAdd(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

// 8. Higher-Order Functions (Returning a Function)
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 9. Methods (Functions with Receivers)
type Rectangle struct {
	width, height int
}

func (r Rectangle) area() int {
	return r.height * r.width
}

// 11. Panic and Recover in Functions
func myPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover")
		}
	}()
	panic("Something went wrong!")
}

func main() {

	// 5. Anonymous (or Lambda) Functions
	newAdd := func(a int, b int) int {
		return a + b
	}

	// 6. Anonymous Functions with Immediate Execution
	immAdd := func(a int, b int) int {
		return a + b
	}(9, 9)

	// first
	firResult := func(a int, b int) int {
		return a + b
	}

	// higher
	double := multiplier(4)(8)
	// or
	// double := multiplier(4)

	// methods
	calcArea := Rectangle{width: 12, height: 12}

	// 10. Defer in Functions
	defer fmt.Println("This is defer and will run at last")

	fmt.Println(add(1, 2))
	fmt.Println(divide(1, 0))
	fmt.Println(swap(1, 20))
	fmt.Println(sum(1, 20, 1, 2, 3, 4, 4, 3, 5, 5))
	fmt.Println(newAdd(1, 12))
	fmt.Println(immAdd)
	fmt.Println(firAdd(4, 4, firResult))
	fmt.Println(double)
	fmt.Println(calcArea.area())
	// or
	// fmt.Println(double(8))
	myPanic()

}
