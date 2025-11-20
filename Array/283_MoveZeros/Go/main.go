package main

import "fmt"

func moveZeros(nums []int) {

	left := 0
	for right := range nums {

		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}

func main() {

	nums := []int{0, 1, 0, 3, 12}
	moveZeros(nums)

	fmt.Println("After moving: ", nums)
}
