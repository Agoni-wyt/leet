package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {

	output := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(output)
}

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意： 答案中不可以包含重复的三元组。
示例：
给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]
*/
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	// 找出a + b + c = 0
	// a = nums[i], b = nums[left], c = nums[right]
	for i := 0; i < len(nums)-2; i++ {
		// 排序之后如果第一个元素已经大于零，那么无论如何组合都不可能凑成三元组，直接返回结果就可以了
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		// 去重a
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 去重逻辑应该放在找到一个三元组之后，对b 和 c去重
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] {
			continue // 重复数字跳过
		}
		if x > 0 { // 连续三个大于0，后面只会更大，没必要再算
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 {
			continue // +最大的两个值都<0,当前值应该继续变大
		}
		j, k := i+1, n-1 // 当前指针i，左指针i+1，右指针n-1
		for j < k {      // 左右指针重合
			s := x + nums[j] + nums[k]
			if s > 0 { // 数太大了，大值要减少，右指针向左
				k--
			} else if s < 0 { // 数太小了，小值向右，左指针向右
				j++
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 左指针向右，如果是重复数字，继续向右
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // 右指针向左，如果是重复数字，继续向左
			}
		}
	}
	return ans
}
