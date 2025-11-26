package main

func mergeNodes(head *Node) *Node {

	if head == nil {
		return nil
	}

	dummy := newNode(0)
	tail := dummy

	current := head.next
	sum := 0

	for current != nil {
		if current.value == 0 {
			if sum > 0 {
				tail.next = newNode(sum)
				tail = tail.next
				sum = 0
			}
		} else {
			sum += current.value
		}
		current = current.next
	}

	return dummy.next
}

func main() {

	l1 := newLinkedList()
	l1.append(newNode(0))
	l1.append(newNode(3))
	l1.append(newNode(1))
	l1.append(newNode(0))
	l1.append(newNode(4))
	l1.append(newNode(5))
	l1.append(newNode(2))
	l1.append(newNode(0))

	l1.head = mergeNodes(l1.head)

	l1.printList()
}
