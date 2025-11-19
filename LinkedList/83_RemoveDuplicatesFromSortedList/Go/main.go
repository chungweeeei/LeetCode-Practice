package main

// 刪除排序鏈表中的重複元素
// 輸入: 已排序的鏈表 head
// 輸出: 刪除重複元素後的鏈表
func deleteDuplicates(head *Node) *Node {

	// 如果鏈表是空的，直接回傳 nil
	if head == nil {
		return nil
	}

	// 建立一個虛擬頭節點，方便處理
	dummy := newNode(0)
	dummy.next = head

	// slow: 指向目前已確認不重複的節點
	// fast: 用來往前探索的指針
	slow := head
	fast := head

	// 用 fast 指針走過整個鏈表
	for fast != nil {
		// 如果 slow 和 fast 的值不同，說明找到了新的不重複值
		if slow.value != fast.value {
			slow.next = fast // 將新的不重複節點連接到結果鏈表
			slow = slow.next // slow 往前移動
		}
		fast = fast.next // fast 繼續往前探索
	}

	// 斷開 slow 後面可能還連著的重複節點
	slow.next = nil

	// 回傳新的鏈表頭（跳過虛擬頭節點）
	return dummy.next
}

func main() {

	l1 := newLinkedList()

	l1.append(newNode(1))
	l1.append(newNode(1))
	l1.append(newNode(2))
	l1.append(newNode(3))
	l1.append(newNode(3))

	l1.printList()

	l1.head = deleteDuplicates(l1.head)

	l1.printList()
}
