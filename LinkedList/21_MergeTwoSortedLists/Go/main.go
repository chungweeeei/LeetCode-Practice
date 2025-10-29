package main

func mergeTwoLists(l1 *Node, l2 *Node) *LinkedList {

	if l1 == nil && l2 == nil {
		return nil
	}

	sortedLinkedList := newLinkedList()
	for l1 != nil || l2 != nil {

		if l1 == nil {
			sortedLinkedList.append(l2)
			l2 = l2.next
			continue
		}

		if l2 == nil {
			sortedLinkedList.append(l1)
			l1 = l1.next
			continue
		}

		if l1.value < l2.value {
			sortedLinkedList.append(l1)
			l1 = l1.next
		} else {
			sortedLinkedList.append(l2)
			l2 = l2.next
		}
	}

	return sortedLinkedList
}

func main() {

	l1 := newLinkedList()
	l1.append(newNode(1))
	l1.append(newNode(2))
	l1.append(newNode(4))

	l2 := newLinkedList()
	l2.append(newNode(1))
	l2.append(newNode(3))
	l2.append(newNode(4))

	mergeList := mergeTwoLists(l1.head, l2.head)
	mergeList.printList()
}
