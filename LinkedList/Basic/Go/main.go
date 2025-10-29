package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func newNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func newLinkedList() *LinkedList {
	return &LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *LinkedList) append(node *Node) {

	if node == nil {
		return
	}

	if l.head == nil {
		l.head = node
		l.tail = l.head
		l.length++
		return
	}

	l.tail.next = node
	l.tail = node
	l.length++
}

func (l *LinkedList) prepend(node *Node) {

	if node == nil {
		return
	}

	if l.head == nil {
		l.head = node
		l.tail = l.head
		l.length++
		return
	}

	node.next = l.head
	l.head = node
	l.length++
}

func (l *LinkedList) insert(node *Node, idx int) {

	if idx <= 0 {
		l.prepend(node)
		return
	}

	if idx > l.length-1 {
		l.append(node)
		return
	}

	prevNode := l.traversalToIndex(idx - 1)

	node.next = prevNode.next
	prevNode.next = node
	l.length++
}

func (l *LinkedList) remove(idx int) {

	if idx <= 0 {
		l.head = l.head.next
		l.length--
		return
	}

	if idx > l.length-1 {
		idx = l.length - 1
	}

	prevNode := l.traversalToIndex(idx - 1)
	prevNode.next = prevNode.next.next
	l.length--
}

func (l *LinkedList) traversalToIndex(idx int) *Node {

	currentNode := l.head
	count := 0
	for currentNode != nil {

		if count == idx {
			return currentNode
		}
		currentNode = currentNode.next
		count++
	}

	return currentNode
}

func (l *LinkedList) printList() {

	result := []int{}
	if l.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	currentNode := l.head
	for currentNode != nil {
		result = append(result, currentNode.value)
		currentNode = currentNode.next
	}

	fmt.Println("Linked List: ", result)
}

func main() {

	l := newLinkedList()
	l.append(newNode(1))
	l.append(newNode(2))
	l.append(newNode(3))
	l.append(newNode(4))
	l.append(newNode(5))

	l.printList()

	l.insert(newNode(6), 1)

	l.printList()

	l.remove(5)

	l.printList()
}
