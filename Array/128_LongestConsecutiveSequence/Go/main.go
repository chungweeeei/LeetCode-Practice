package main

import "fmt"

func longestConsecutive(nums []int) int {

	hash := make(map[int]struct{})
	for _, num := range nums {
		hash[num] = struct{}{}
	}

	longest := 0

	for _, num := range nums {

		if _, existed := hash[num+1]; existed {

			currentNum := num
			currentStreak := 1

			for {

				if _, existed := hash[currentNum+1]; !existed {
					break
				}

				currentNum += 1
				currentStreak += 1
			}

			longest = max(longest, currentStreak)
		}
	}

	return longest
}

func main() {

	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))

}
