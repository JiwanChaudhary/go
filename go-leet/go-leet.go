package main

import (
	"fmt"
	"strconv"
)

// func isPalindrome(x int) bool {
// 	original := strconv.Itoa(x)
// 	runes := []rune(original)
// 	var reversed string

// 	fmt.Println(runes, "runes")
// 	fmt.Println(len(runes), "runesLen")

// 	for i := len(strconv.Itoa(x)) - 1; i >= 0; i-- {
// 		reversed += string(original[i])
// 	}

// 	return original == reversed

// }

func isPalindrome(x int) bool {
	original := strconv.Itoa(x)
	runes := []rune(original)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(1234321))
}
