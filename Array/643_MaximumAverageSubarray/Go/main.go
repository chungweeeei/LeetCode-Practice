package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {

	sum := 0

	// 1. 先計算第一個視窗 (前 k 個元素) 的總和
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	// 初始化最大總和為第一個視窗的總和
	maxSum := sum

	// 2. 開始滑動視窗，從第 k 個元素開始往後遍歷
	for i := k; i < len(nums); i++ {
		// 核心邏輯：加上新進來的 (nums[i])，減去最左邊離開的 (nums[i-k])
		sum += nums[i] - nums[i-k]

		// 更新最大總和
		if sum > maxSum {
			maxSum = sum
		}
	}

	// 3. 計算平均值並回傳
	return float64(maxSum) / float64(k)
}

func main() {

	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Printf("Max Average: %v\n", findMaxAverage(nums, k)) // 預期輸出: 12.75
}
