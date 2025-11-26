package main

func mergeBetween(node1 *Node, a, b int, node2 *Node) *Node {

	prevNode := node1
	for i := 0; i < a-1; i++ {
		prevNode = prevNode.next
	}

	endNode := prevNode
	for i := 0; i < (b - a + 1); i++ {
		endNode = endNode.next
	}

	afterNode := endNode.next
	prevNode.next = node2

	tail2 := node2
	for tail2.next != nil {
		tail2 = tail2.next
	}

	tail2.next = afterNode

	return node1
}

func main() {

	l1 := newLinkedList()
	l1.append(newNode(10))
	l1.append(newNode(1))
	l1.append(newNode(13))
	l1.append(newNode(6))
	l1.append(newNode(9))
	l1.append(newNode(5))

	l2 := newLinkedList()
	l2.append(newNode(1000000))
	l2.append(newNode(1000001))
	l2.append(newNode(1000002))

	mergeBetween(l1.head, 3, 4, l2.head)

	l1.printList()
}
