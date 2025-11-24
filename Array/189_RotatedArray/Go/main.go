package main

import "fmt"

func reverse(nums []int, start int, end int) {

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}

}

func rotate(nums []int, k int) {

	n := len(nums)
	k = k % n

	// reverse whole array
	reverse(nums, 0, n-1)

	// reverse forward
	reverse(nums, 0, k-1)

	// reverse backward
	reverse(nums, k, n-1)
}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println("Rotate Array: ", nums)

	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println("Rotate Array: ", nums)
}
