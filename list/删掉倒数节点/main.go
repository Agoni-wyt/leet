package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

/*
删除倒数第n个节点，可以使用快慢指针，让快慢指针的插值为n，，当快指针到指到末尾时，慢指针所指的就是要被删除的节点
输入：head = [1,2,3,4,5], n = 2 输出：[1,2,3,5] 示例 2：

输入：head = [1], n = 1 输出：[] 示例 3：

输入：head = [1,2], n = 1 输出：[1]

fast要先走n + 1步 ，因为要将快指针指到空时，慢指针刚好指到目标节点的上一个，以便成功进行删除操作
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 设置虚拟头
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead
	// 快指针先走n+1
	for fast != nil && n >= 0 {
		fast = fast.Next
		n--
	}
	// 快慢指针同时移动
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 删除当前慢指针的下一个
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

// 也可以算出整个列表的长度length，然后找到length-n-1位置的节点，然后删除改节点的next
