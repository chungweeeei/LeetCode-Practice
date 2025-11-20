package main

import "fmt"

func removeElement(nums []int, val int) int {

	left := 0
	for right := range nums {

		if nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	fmt.Println("After removing:", nums)
	return left
}

func main() {

	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))

}
