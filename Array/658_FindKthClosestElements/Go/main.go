package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {

	low := 0
	high := len(arr) - 1

	for high-low >= k {

		if x-arr[low] > arr[high]-x {
			low++
		} else {
			high--
		}
	}

	return arr[low : high+1]
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(arr, 4, 3))

	arr = []int{1, 1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(arr, 4, -1))
}
