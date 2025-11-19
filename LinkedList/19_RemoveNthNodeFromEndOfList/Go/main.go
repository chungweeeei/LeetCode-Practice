package main

/*
核心技巧：雙指針法 (Two Pointers)
- 讓兩個指針保持 n+1 的距離
- 當前面的指針到達末尾時，後面的指針正好在要刪除節點的前一個位置

時間複雜度：O(L) - L 是鏈表長度，只需要遍歷一次
空間複雜度：O(1) - 只使用了常數額外空間
*/

func removeNthFromEnd(head *Node, n int) *Node {

	if head == nil {
		return nil
	}

	// 為什麼需要建立虛擬節點？
	// 因為要刪除的可能是頭節點，虛擬節點可以統一處理邏輯
	// 例如：[1,2], n=2 要刪除節點1，有了dummy就不需要特殊處理
	dummy := newNode(0)
	dummy.next = head

	// 第一步：讓 first 指針先走 n+1 步
	// 為什麼是 n+1？因為我們要讓 second 指針停在要刪除節點的「前一個」位置
	first := dummy
	second := dummy
	for i := 0; i < n+1; i++ {
		if first != nil {
			first = first.next
		}
	}

	// 第二步：兩個指針同時向前移動
	for first != nil {
		first = first.next
		second = second.next
	}

	// 第三步：刪除節點 (將下個Node指向下下個Node)
	second.next = second.next.next

	return dummy.next
}

func main() {

	l1 := newLinkedList()
	l1.append(newNode(1))
	l1.append(newNode(2))
	l1.append(newNode(3))
	l1.append(newNode(4))
	l1.append(newNode(5))

	l1.head = removeNthFromEnd(l1.head, 2)

	l1.printList()
}
