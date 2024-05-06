package main

import "fmt"

func main() {
	input := 133
	output := isHappy(input)
	fmt.Println(output)
}

/*
编写一个算法来判断一个数 n 是不是快乐数。
「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。
如果 n 是快乐数就返回 True ；不是，则返回 False 。
示例：
输入：19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

题目中说了会 无限循环，那么也就是说求和的过程中，sum会重复出现，这对解题很重要！

正如：关于哈希表，你该了解这些！ (opens new window)中所说，当我们遇到了要快速判断一个元素是否出现集合里的时候，就要考虑哈希法了。
*/
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	fmt.Println("sum: ", sum)
	return sum

}

/*
等价于链表是否有环，使用快慢指针，
初始，slow，fast:=n,step(n)
慢指针每次走1步，也就是step(slow)
块指针每次两步，也就是step(step(fast))
当快指针不等于1和慢指针时，快慢指针一直走下一步
循环结束时判断，快指针是否==1，如果等于则证明有环
*/

func isHappy2(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
