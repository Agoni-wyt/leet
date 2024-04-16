package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 双指针
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
}
