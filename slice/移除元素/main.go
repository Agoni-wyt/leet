package main

import (
	"fmt"
)

func main() {
	q5()

}

/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。 你不需要考虑数组中超出新长度后面的元素。
示例 2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
*/
func q1() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	output := fastAnsSlowPoint(nums, val)
	output = directionPointer(nums, val)
	fmt.Println(output)
}
func fastAnsSlowPoint(nums []int, val int) int {
	fast, slow, n := 0, 0, len(nums)
	for fast < n {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
func directionPointer(nums []int, val int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == val {
			nums[l] = nums[r]
			r--
		} else {
			l++
		}

	}
	return l
}

/*
一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
func q2() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	output := removeDuplicates(nums)
	fmt.Println(output)
}
func removeDuplicates(nums []int) int {
	s, f, n := 0, 1, len(nums)
	for f < n {
		if nums[f] != nums[s] {
			s++
			nums[s] = nums[f]
		}
		f++
	}
	fmt.Println(nums)

	return s + 1
}

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。
示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:
输入: nums = [0]
输出: [0]
*/
func q3() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	output := moveZeroes2(nums)
	fmt.Println(output)
}
func moveZeroes(nums []int) []int {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}
func moveZeroes2(nums []int) []int {
	s, f, n := 0, 0, len(nums)
	for f < n {
		// 找到不是0的数，将其换到前面来，换完了s++
		if nums[f] != 0 {
			nums[s], nums[f] = nums[f], nums[s]
			s++
		}
		f++
	}
	return nums
}

/*
给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
注意：如果对空文本输入退格字符，文本继续为空。
示例 1：
输入：s = "ab#c", t = "ad#c"
输出：true
解释：s 和 t 都会变成 "ac"。
示例 2：
输入：s = "ab##", t = "c#d#"
输出：true
解释：s 和 t 都会变成 ""。
示例 3：
输入：s = "a#c", t = "b"
输出：false
解释：s 会变成 "c"，但 t 仍然是 "b"。
*/
func q4() {
	s := "a##c"
	t := "#a#c"
	output := backspaceCompare(s, t)
	fmt.Println(output)
}
func backspaceCompare(s string, t string) bool {
	s = exec(s) // Com(s)
	t = exec(t) // Com(t)
	return s == t
}
func exec(str string) string {
	s := []byte{}
	for i := range str {
		if str[i] != '#' {
			s = append(s, str[i])
		} else if len(s) > 0 {
			s = s[:len(s)-1]
		}
	}
	return string(s)
}

// Com 栈移除方法
func Com(s string) string {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			stack = append(stack, s[i])
		} else if len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
func twoPoint(s, t string) bool {
	// 因为该题，#对右边字符没影响，所以应该才右边计算
	// 应该逆序考虑,两个字符串对应指针位置是否相同
	// 遇到 # 就跳过前面几个
	skipS, skipT := 0, 0       // 存放字符串中#的数量
	i, j := len(s)-1, len(t)-1 // 指针
	for i >= 0 || j >= 0 {
		//s 字符串当前指针
		for i >= 0 {
			if s[i] == '#' {
				skipS++ //若当前字符是 # ，则 skipS自增 1；-->记入即将要消除的字符数
				i--
			} else if skipS > 0 { //若当前字符不是 # ，且skipS不为0，则 skipS自减 1；-->该字符被消除
				skipS--
				i--
			} else { // 消除数为0，且不是#，则保留，不需要消除
				break
			}
		}
		//t 字符串当前指针
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] { // 任意当前字符不相同时，遍历结束
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}

		i--
		j--
	}
	return true
}
func backspaceCompare2(s string, t string) bool {
	ss := []byte{}
	st := []byte{}
	// 遇到#，将字符串末尾的值移除，没遇到#，就存到字符里
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(ss) > 0 {
				ss = ss[:len(ss)-1]
			}
		} else {
			ss = append(ss, s[i])
		}
	}

	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, t[i])
		}
	}
	return string(ss) == string(st)
}

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/
func q5() {
	var input = []int{-4, -1, 0, 3, 10}
	output := sortedSquares(input)
	fmt.Println(output)
}
func sortedSquares(nums []int) []int {
	// 离0越近，平方后越小
	// 两种方法，1. 平方后排序（或者绝对值后排序然后平方）
	//2. 双指针，从倒序来，找出两端最大的放入新数组末尾
	l, r, n := 0, len(nums)-1, len(nums)-1
	ans := make([]int, len(nums))
	for l <= r {
		left := nums[l] * nums[l]
		right := nums[r] * nums[r]
		if left > right {
			ans[n] = left
			l++
		} else {
			ans[n] = right
			r--
		}
		n--
	}
	return ans
}
