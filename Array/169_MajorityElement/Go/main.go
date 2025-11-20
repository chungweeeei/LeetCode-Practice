package main

import "fmt"

func majorityElement(nums []int) int {

	/*
		Boyer-Moore Voting Algorithm

		use candidate variable to store the current candidate for majority element
		and a count variable to track the count of the candidate.

		iterate through each number in the array:
			if count is 0, set candidate to the current number and increment count
			else if the current number is equal to candidate, increment count
			else decrement count

		at the end of the iteration, candidate will be the majority element
	*/

	candidate := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
			count++
		} else if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {

	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
