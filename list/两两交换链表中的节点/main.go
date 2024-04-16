package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	/*
		链表两两翻转

		1->2->3->4->5->6

		2->1->4->3->6->5

		输入：[1,2,3,4]
		输出：[2,1,4,3]

		输入：[1]
		输出：[1]
	*/

}
func swapPairs(head *ListNode) *ListNode {
	// 使用假头
	mockHead := &ListNode{Next: head}
	pre := mockHead
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		//pre=list[(i+2)-1]
		pre = head
		//head=list[(i+2)]
		head = next
	}
	return mockHead.Next
}
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
