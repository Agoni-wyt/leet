package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

/*
思路
https://programmercarl.com/0142.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8II.html#%E6%80%9D%E8%B7%AF
*/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 找到环内相同节点
			for slow != head { // 变为一个一个移动
				slow = slow.Next // 从当前位置移动
				head = head.Next // 从头移动
			}
			return head // 找到从头移动的节点
		}
	}
	return nil
}

// 哈希表方法，遇见的都存入字典，找到存在的，返回当前节点
func detectCycle2(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}
