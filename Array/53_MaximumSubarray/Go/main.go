package main

import "fmt"

func maxSubarray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	current := nums[0]
	maxSum := nums[0]

	for _, num := range nums {
		current = max(num, current+num)
		maxSum = max(current, maxSum)
	}

	return maxSum
}

func main() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Max Sum:", maxSubarray(nums))

	nums = []int{1}
	fmt.Println("Max Sum:", maxSubarray(nums))
}
