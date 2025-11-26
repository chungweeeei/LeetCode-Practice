package main

import "fmt"

func generate(numsRow int) [][]int {

	if numsRow == 0 {
		return [][]int{}
	} else if numsRow == 1 {
		return [][]int{{1}}
	}

	result := make([][]int, numsRow)
	result[0] = []int{1}

	for i := 1; i < numsRow; i++ {
		result[i] = append(result[i], 1)

		for j := 1; j < len(result[i-1]); j++ {
			result[i] = append(result[i], result[i-1][j]+result[i-1][j-1])
		}

		result[i] = append(result[i], 1)
	}

	return result
}

func main() {

	fmt.Println(generate(0))
	fmt.Println(generate(1))
	fmt.Println(generate(5))

}
