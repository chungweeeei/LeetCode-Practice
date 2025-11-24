package main

import (
	"sort"
)

func ThreeSum(nums []int) {

	// sorted
	sort.Ints(nums)
	// nums => [-4, -1, -1, 0, 1, 2]
	for i := 0; i < len(nums)-2; i++ {

		// remove duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

	}
}

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	ThreeSum(nums)
}
