package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

/*
找到两个列表开始相交的起始节点
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

两个链表在节点 c1(指针) 开始相交：

	a1->a2->c1->c2->c3

b1->b2->b3->c1->c2->c3
题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。

输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'
解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
这两个链表不相交，因此返回 null 。

双指针思想，可以计算长度，然后将长链表上的指针移动到倒数短链表长度的位置，再开始比较，即计算链表路径差

也可以使用数学思想，将总路程一定为 len(a)+len(b)
其中a链表中不同部分设为diffA，b链表中不同部分设为diffB,俩链表相同部分为C
即len(a) = diffA+c ; len(b) = diffB+c
a:  ---**   diffA=3  c=2
b:  =**		diffB=1  c=2
总： len(a)+len(b) = diffA+diffB+c+c = 8
A指针路线   ---**=**
B指针路线	  =**---**
依据上图当两指针都走了 diffA+diffB+c时，指针位置一定对应，因为他们剩下的路程皆为c

当两链表没有相交时
x:   ----   diffX = 4
y:   ==		diffY = 2
总：  len(x)+len(y) = diffX+diffY = 6
X指针路线   ----==
Y指针路线	   ==----
依据上图两指针都走了 diffX+diffY 时，指针位置皆为空值

要额外注意，当两链表任意为空时，一定不存在相交节点，直接返回空
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var fast, slow *ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
