package main

import "fmt"

func addTwoNumbers(left *Node, right *Node) []int {

	result := []int{}

	addOne := 0
	for left != nil || right != nil {

		if left == nil {
			sum := right.value + addOne
			addOne = sum / 10
			sum = sum % 10

			result = append(result, sum)
			right = right.next
			continue
		}

		if right == nil {
			sum := left.value + addOne
			addOne = sum / 10
			sum = sum % 10

			result = append(result, sum)
			left = left.next
			continue
		}

		sum := left.value + right.value + addOne
		addOne = sum / 10
		sum = sum % 10

		result = append(result, sum)
		left = left.next
		right = right.next
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
