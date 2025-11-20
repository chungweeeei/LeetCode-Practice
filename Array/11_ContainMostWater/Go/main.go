package main

import "fmt"

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1

	maxArea := 0
	for left < right {

		// area = width * height
		area := (right - left) * min(height[left], height[right])
		maxArea = max(area, maxArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))

	height = []int{1, 1}
	fmt.Println(maxArea(height))

}
