package main

/*
核心思路：
1. 使用虛擬頭節點簡化邊界處理
2. 每次處理一對節點，調整三個連接
3. 移動指針到下一對節點

時間複雜度：O(n) - 遍歷整個鏈表一次
空間複雜度：O(1) - 只使用常數額外空間
*/

func swapPairs(head *Node) *Node {

	if head == nil {
		return nil
	}

	dummy := newNode(0)
	dummy.next = head

	// prev：指向當前處理的一對節點的前一個節點
	// 初始時指向虛擬節點，這樣第一對節點的「前一個」就是 dummy
	prev := dummy
	for head != nil && head.next != nil {

		// 標記要交換的兩個節點
		first := head
		second := head.next

		/*
		   交換前的狀態：
		   prev -> first -> second -> (其他節點)

		   我們要變成：
		   prev -> second -> first -> (其他節點)

		   需要修改三個連接：
		   1. prev.next = second （prev 指向 second）
		   2. first.next = second.next （first 指向 second 原來的下一個）
		   3. second.next = first （second 指向 first）
		*/
		prev.next = second
		first.next = second.next
		second.next = first

		/*
		   交換完成後：
		   prev -> second -> first -> (其他節點)
		*/

		// 更新指針，準備處理下一對節點
		// prev 移動到當前這對的最後一個節點（交換後的 first）
		prev = first

		// head 移動到下一對的第一個節點
		head = first.next
	}

	return dummy.next
}

func main() {

	l1 := newLinkedList()
	l1.append(newNode(1))
	l1.append(newNode(2))
	l1.append(newNode(3))
	l1.append(newNode(4))

	l1.head = swapPairs(l1.head)

	l1.printList()
}
