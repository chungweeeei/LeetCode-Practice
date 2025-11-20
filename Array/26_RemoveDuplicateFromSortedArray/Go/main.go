package main

import "fmt"

func removeDuplicates(nums []int) int {

	left := 0
	for right := range nums {

		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}

	return left + 1
}

func main() {

	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))

}
