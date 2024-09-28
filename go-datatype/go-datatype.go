// Basic data types
// 1. Boolean
// 2. Numeric
// 3. String

// Composite data types
// 4. Array
// 5. Slices
// 6. Maps
// 7. Structs
// 8. Pointers

package main

import "fmt"

func main() {
	var isActive bool = false // boolean
	var num int = 12          // numeric
	var name string = "Jiwan" // string

	var arr [5]int = [5]int{1, 2, 3, 4, 5}                                       // array -> fixed size
	var slice []int = []int{1, 2, 3, 4, 5, 5, 6}                                 // slices -> dynamic size
	var maps map[string]string = map[string]string{"name": "Jiwan", "age": "12"} // map

	type Person struct { // structs
		name string
		age  int
	}

	p := Person{
		name: "Jiwan",
		age:  13,
	}

	var pointStruct *string = &p.name

	a := 10 // pointer
	var point *int = &a

	tNum := 1398977767655665086

	tFloat := 1.22455242545254254242535354354345 // float64

	fmt.Println(isActive, num, name, arr, slice, len(slice), maps, len(maps), p, p.age, pointStruct, point, tFloat, tNum)
}
