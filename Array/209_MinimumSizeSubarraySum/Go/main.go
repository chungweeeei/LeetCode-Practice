package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {

	left := 0
	sum := 0
	minLen := len(nums) + 1

	for right := range nums {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}

			sum = sum - nums[left]
			left += 1
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}

func main() {

	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}
