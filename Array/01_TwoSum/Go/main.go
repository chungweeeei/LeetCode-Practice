package main

import "fmt"

func twoSum(nums []int, target int) []int {

	hash := map[int]int{}

	for i, num := range nums {

		diff := target - num
		if index, existed := hash[diff]; existed {
			return []int{index, i}
		}

		hash[num] = i
	}

	return []int{}
}

func main() {

	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))

	nums = []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))

	nums = []int{3, 3}
	fmt.Println(twoSum(nums, 6))
}
