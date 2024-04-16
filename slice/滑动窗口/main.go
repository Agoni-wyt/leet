package main

import (
	"fmt"
)

func main() {

}

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：
输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/
func q1() {
	input := []int{2, 3, 1, 2, 4, 3}
	output := minSubArrayLen(7, input)
	fmt.Println(output)
}
func minSubArrayLen(target int, nums []int) int {
	// 从前向后，先移动右边界，当其中内容>目标值后，持续移动左边界，缩短窗口，然后再移动右边界，滑动窗口
	ans := len(nums) + 1
	sum := 0
	start := 0
	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
	}
	if ans > len(nums) {
		return 0
	} else {
		return ans
	}
}

/*
你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i] 是第 i 棵树上的水果 种类 。

你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：

你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合篮子中的水果类型。每采摘一次，你将会向右移动到下一棵树，并继续采摘。
一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。

示例 1：

输入：fruits = [1,2,1]
输出：3
解释：可以采摘全部 3 棵树。
示例 2：

输入：fruits = [0,1,2,2]
输出：3
解释：可以采摘 [1,2,2] 这三棵树。
如果从第一棵树开始采摘，则只能采摘 [0,1] 这两棵树。
示例 3：

输入：fruits = [1,2,3,2,2]
输出：4
解释：可以采摘 [2,3,2,2] 这四棵树。
如果从第一棵树开始采摘，则只能采摘 [1,2] 这两棵树。
示例 4：

输入：fruits = [3,3,3,1,2,1,1,2,3,3,4]
输出：5
解释：可以采摘 [1,2,1,1,2] 这五棵树。
*/
func q2() {
	input := []int{2, 3, 1, 2, 4, 3}
	output := totalFruit(input)
	fmt.Println(output)
}
func totalFruit(fruits []int) int {
	// 保证两个总类，当第三个总类出现时，移动start，使其仅保留两个总类
	// 总数为 end-start+1 0~len(fruits)
	basket := make(map[int]int)
	start, ans := 0, 0
	for end, f1 := range fruits {
		basket[f1]++          // 放入篮子
		for len(basket) > 2 { // 发现篮子中已经有两种了
			f2 := fruits[start]
			basket[f2]-- // 把地下先放进去发水果拿出去，-个一个拿出去
			if basket[f2] == 0 {
				delete(basket, f2)
			}
			start++
		}
		ans = max(ans, end-start+1)
	}
	return ans
}
func totalFruit2(fruits []int) (ret int) {
	first, second, lastIdx, left := 0, 0, 0, 0
	for right := 0; right < len(fruits); right++ {
		if fruits[right] != first && fruits[right] != second {
			first, second, left = fruits[lastIdx], fruits[right], lastIdx
		}
		if fruits[right] != fruits[lastIdx] {
			lastIdx = right
		}
		ret = max(ret, right-left+1)
	}
	return
}

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
*/
func q3() {
	input := "ADOBECODEBANC"
	t := "BANC"
	output := minWindow(input, t)
	fmt.Println(output)
}
func minWindow(s string, t string) string {
	tMp, sMp := map[byte]int{}, map[byte]int{}
	ansL, ansR := -1, len(s)
	left := 0
	for i := range t {
		tMp[t[i]]++
	}
	for right := range s { // 移动子串右端点
		sMp[s[right]]++       // 右端点字母移入子串
		for check(tMp, sMp) { // 涵盖
			if right-left < ansR-ansL { // 找到更短的子串
				ansL, ansR = left, right // 记录此时的左右端点
			}
			sMp[s[left]]-- // 左端点字母移出子串
			left++         // 移动子串左端点
		}
	}
	if ansL < 0 {
		return ""
	}
	return s[ansL : ansR+1]
}
func check(ori, cnt map[byte]int) bool {
	//检查当前mp是否包含orimp
	for k, v := range ori {
		if cnt[k] < v {
			return false
		}
	}
	return true
}
