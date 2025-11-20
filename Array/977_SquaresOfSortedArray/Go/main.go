package main

import "fmt"

func sortedSquares(nums []int) []int {

	left := 0
	right := len(nums) - 1

	result := make([]int, len(nums))
	pos := len(nums) - 1

	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			result[pos] = nums[left] * nums[left]
			left++
			pos--
		} else {
			result[pos] = nums[right] * nums[right]
			right--
			pos--
		}
	}

	return result
}

func main() {

	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))

	nums = []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}
