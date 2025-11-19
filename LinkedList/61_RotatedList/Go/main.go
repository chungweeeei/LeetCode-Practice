package main

/*
核心思路：
1. 計算鏈表長度，找到尾節點
2. 將鏈表形成環形
3. 找到新的尾節點和頭節點
4. 斷開環形，形成新的鏈表

時間複雜度：O(n) - 需要遍歷鏈表計算長度
空間複雜度：O(1) - 只使用常數額外空間
*/

func rotateRight(head *Node, k int) *Node {

	if head == nil || head.next == nil {
		return nil
	}

	// 第一步：計算鏈表長度，同時找到尾節點
	length := 1
	tail := head
	for tail.next != nil {
		length++
		tail = tail.next
	}

	// 迴圈結束後，tail 指向最後一個節點，length 是鏈表總長度
	// 第二步：優化 k 值
	// 如果 k 是長度的倍數，旋轉後鏈表不變
	// 例如：長度5的鏈表，k=5 等於 k=0，k=7 等於 k=2
	k = k % length
	if k == 0 {
		return head
	}

	// 第三步：找到旋轉後的新尾節點
	// 新尾節點的位置 = length - k - 1
	// 為什麼減1？因為我們要找的是節點位置（從0開始計算）
	/*
	   例子：[1,2,3,4,5], k=2, length=5
	   新尾節點位置 = 5-2-1 = 2（即節點3，索引從0開始）
	   結果：4 -> 5 -> 1 -> 2 -> 3
	             新頭        新尾
	*/
	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.next
	}

	// 第四步：確定新的頭節點
	// 新頭節點就是新尾節點的下一個
	newHead := newTail.next

	// 第五步：斷開鏈表
	// 在新尾節點處斷開，形成新的鏈表尾部
	newTail.next = nil

	// 第六步：連接成新的鏈表
	// 將原來的尾節點連接到原來的頭節點
	// 這樣就把後面的部分移到了前面
	tail.next = head

	return newHead
}

func main() {

	l1 := newLinkedList()
	l1.append(newNode(1))
	l1.append(newNode(2))
	l1.append(newNode(3))
	l1.append(newNode(4))
	l1.append(newNode(5))

	l1.head = rotateRight(l1.head, 2)

	l1.printList()
}
