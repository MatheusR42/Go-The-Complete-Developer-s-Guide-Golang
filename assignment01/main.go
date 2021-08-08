package main

import "fmt"

func createNumList() []int {
	nums := []int{}

	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
	}

	return nums
}

func main() {
	nums := createNumList()

	for _, num := range nums {
		if num % 2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}