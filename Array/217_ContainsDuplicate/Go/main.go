package main

import "fmt"

func containDuplicate(nums []int) bool {

	// 使用 struct{} 空結構體，不佔用額外記憶體
	hash := make(map[int]struct{})

	for _, num := range nums {

		if _, existed := hash[num]; existed {
			return true
		}

		hash[num] = struct{}{}
	}

	return false
}

func main() {

	nums := []int{1, 2, 3, 1}
	fmt.Println(containDuplicate(nums))

	nums = []int{1, 2, 3, 4}
	fmt.Println(containDuplicate(nums))

}
