package main

import "fmt"

/*
時間複雜度：O(max(m,n)) - m,n 分別是兩個鏈表的長度
空間複雜度：O(max(m,n)) - 結果鏈表的長度
*/

func addTwoNumbers(left *Node, right *Node) []int {

	result := []int{}

	// 進位值：當兩位數相加 >= 10 時，需要向下一位進位
	// 例如：7 + 8 = 15，當前位存 5，進位值是 1
	addOne := 0

	for left != nil || right != nil {

		// 情況1：左鏈表已經走完，只剩右鏈表
		if left == nil {
			sum := right.value + addOne
			addOne = sum / 10
			sum = sum % 10

			result = append(result, sum)
			right = right.next
			continue
		}

		// 情況2：右鏈表已經走完，只剩左鏈表
		if right == nil {
			sum := left.value + addOne
			addOne = sum / 10
			sum = sum % 10

			result = append(result, sum)
			left = left.next
			continue
		}

		// 情況3：兩個鏈表都還有節點，正常相加
		// 當前位 = 左節點值 + 右節點值 + 之前的進位
		sum := left.value + right.value + addOne
		addOne = sum / 10
		sum = sum % 10

		result = append(result, sum)

		// 兩個指針都向前移動到下一個節點
		left = left.next
		right = right.next
	}

	// 重要：檢查最後是否還有進位
	// 例如：999 + 1 = 1000，最後還會有一個進位 1
	if addOne > 0 {
		result = append(result, addOne)
	}

	return result
}

func main() {

	l1 := newLinkedList()
	l1.append(newNode(2))
	l1.append(newNode(4))
	l1.append(newNode(3))

	l2 := newLinkedList()
	l2.append(newNode(5))
	l2.append(newNode(6))
	l2.append(newNode(4))

	fmt.Println(addTwoNumbers(l1.head, l2.head))
}
