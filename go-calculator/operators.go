package main

import (
	"fmt"
)

type Variables struct {
	var1 int
	var2 int
}

// func (userInput *Variables) validateInput() error {
// 	fmt.Println("Enter first input:")
// 	fmt.Scan(&userInput.var1)
// 	fmt.Println("Enter second input:")
// 	fmt.Scan(&userInput.var2)

// 	if userInput.var2 == 0 {
// 		err := errors.New("cannot divide by 0")
// 		fmt.Println(err)
// 		return err
// 	}
// 	return nil
// }

func readUserInput(userInput *Variables) {
	fmt.Println("Enter first input:")
	fmt.Scan(&userInput.var1)
	fmt.Println("Enter second input:")
	fmt.Scan(&userInput.var2)
}

func (userInput *Variables) add() {
	readUserInput(userInput)
	fmt.Println("The addition is:", userInput.var1+userInput.var2)
}

func (userInput *Variables) subtract() {
	readUserInput(userInput)
	fmt.Println("The subtraction is:", userInput.var1-userInput.var2)
}

func (userInput *Variables) multiply() {
	readUserInput(userInput)
	fmt.Println("The multiplication is:", userInput.var1*userInput.var2)
}

func (userInput *Variables) division() {
	readUserInput(userInput)
	if userInput.var2 == 0 {
		fmt.Println("can not divide by 0")
		return
	}
	fmt.Println("The division is:", userInput.var1/userInput.var2)
}
