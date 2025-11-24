package main

import "fmt"

func search(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// --- 判斷哪一半是有序的 ---
		// Case 1: 左半邊 [left...mid] 是有序的
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Case 2: 右半邊 [mid...right] 是有序的
			// 判斷 target 是否落在右半邊的範圍內
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("Search Index:", search(nums, 0))

}
