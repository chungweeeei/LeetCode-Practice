package main

import "fmt"

func detectCycle(head *Node) *Node {

	if head == nil {
		return nil
	}

	slow := head
	fast := head
	for fast != nil {
		if slow == fast {
			// Find the Cycle Node
			slow = head
			for slow != fast {
				slow = slow.next
				fast = fast.next
			}
			return slow
		}

		slow = slow.next
		fast = fast.next.next
	}

	return nil
}

func main() {
	l1 := newLinkedList()

	cycleNode := newNode(2)
	l1.append(newNode(3))
	l1.append(cycleNode)
	l1.append(newNode(0))
	l1.append(newNode(-4))

	l1.tail.next = cycleNode

	fmt.Println("The Linked List cycle at:", detectCycle(l1.head).value)
}
