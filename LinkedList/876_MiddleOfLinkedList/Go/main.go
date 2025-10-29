package main

import "fmt"

/*
時間複雜度: O(n) - 只需遍歷一次
空間複雜度: O(1) - 只用了兩個指針, 不需要額外空間
*/
func middleNode(head *Node) *Node {

	if head == nil {
		return head
	}

	// Two Pointer Method
	// 慢指針: 每次走一步
	// 快指針: 每次走兩步
	slow := head
	fast := head

	// 當快指針走完全程時，慢指針走了一半距離，正好在中間.
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func main() {

	l := newLinkedList()
	l.append(newNode(1))
	l.append(newNode(2))
	l.append(newNode(3))
	l.append(newNode(4))
	l.append(newNode(5))

	l.printList()

	middle := middleNode(l.head)

	fmt.Println("Middle Node: ", middle.value)
}
