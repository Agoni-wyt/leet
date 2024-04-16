package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
题意：删除链表中等于给定值 val 的所有节点。
示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]
示例 2： 输入：head = [], val = 1 输出：[]
示例 3： 输入：head = [7,7,7,7], val = 7 输出：[]
*/

func main() {
	// removeElements()
	// removeElements2()
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head

}
func removeElements2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{} //假装一个头
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
