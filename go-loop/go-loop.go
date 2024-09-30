// loop in go
// format
// for initialization; condition; post {
// }

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[count] != nums[i] {
			count++
			nums[count] = nums[i]
		}
	}
	return count + 1
}

func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func main() {
	// basic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("basic for ends")

	// for loop as while loop
	i := 10
	for i < 15 {
		fmt.Println(i)
		i++
	}
	fmt.Println("for as while ends")

	// infinite for loop
	j := 20
	for {
		fmt.Println(j + 1)
		j++
		if j >= 25 {
			break
		}
	}

	fmt.Println("infinite for loop ends")

	// range based for loop
	nums := []int{1, 2, 3, 4, 5}
	for _, value := range nums {
		fmt.Println(value*2, "value") // only value
	}

	for k := range nums {
		if k == 3 {
			continue
		}
		fmt.Println(k, "key")
	}

	fmt.Println("range based for loop ends")
	// println(test())

	fmt.Println(removeElement([]int{3, 2, 2, 2}, 3))
	fmt.Println(string(123))

}
