package main

import "fmt"

func hasCycle(head *Node) bool {

	if head == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil {

		if slow == fast {
			return true
		}

		slow = slow.next
		fast = fast.next.next
	}

	return false

}

func main() {

	l1 := newLinkedList()

	cycleNode := newNode(2)
	l1.append(newNode(3))
	l1.append(cycleNode)
	l1.append(newNode(0))
	l1.append(newNode(-4))

	l1.tail.next = cycleNode

	fmt.Println("Linked List has cycle:", hasCycle(l1.head))

}
