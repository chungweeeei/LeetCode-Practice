package main

func reorderList(head *Node) {

	if head == nil || head.next == nil {
		return
	}

	// find the middle
	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	mid := slow.next

	// cut the list
	slow.next = nil

	// reverse the list
	var prevNode *Node = nil
	currentNode := mid
	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = prevNode

		prevNode = currentNode
		currentNode = nextNode
	}

	// merge two list
	forwardHead := head
	backwardHead := prevNode

	for forwardHead != nil && backwardHead != nil {

		nextForward := forwardHead.next
		nextBackward := backwardHead.next

		forwardHead.next = backwardHead
		backwardHead.next = nextForward

		forwardHead = nextForward
		backwardHead = nextBackward
	}
}

func main() {

	l1 := newLinkedList()
	l1.append(newNode(1))
	l1.append(newNode(2))
	l1.append(newNode(3))
	l1.append(newNode(4))
	l1.append(newNode(5))

	l1.printList()

	reorderList(l1.head)

	l1.printList()
}
