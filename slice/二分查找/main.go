package main

import (
	"fmt"
	"math"
)

func main() {
	//q1()
	q4()
}

/*
例一：
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
-----
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
-----
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -2
解释: 2 不存在 nums 中因此返回 应在的位置为2，
*/
func q1() {
	//nums := []int{-1, 0, 3, 5, 9, 12}
	//target := 9
	nums2 := []int{1, 3, 5, 6}
	target2 := 0

	output := TwoFind(nums2, target2)

	fmt.Println(output)
}
func TwoFind(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		// 右移朝左相当于 除以2的指定次数，并向下取整
		middle := (right-left)>>1 + left
		if target < nums[middle] {
			right = middle - 1
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			return middle
		}

	}
	return right + 1
}

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
进阶：你可以设计并实现时间复杂度为 $O(\log n)$ 的算法解决此问题吗？
示例 :
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
*/
func q2() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	output := FindByTwo(nums, target)
	fmt.Println(output)
}
func FindByTwo(nums []int, tar int) []int {
	// way1
	res := getBorder(nums, tar)
	// way2
	res = searchRange(nums, tar)
	return res
}
func getBorder(nums []int, tar int) []int {
	left, right := 0, len(nums)-1
	lb, rb := -1, -1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > tar {
			right = mid - 1
		} else if nums[mid] < tar {
			left = mid + 1
		} else {
			if nums[left] != tar {
				left++
			}
			if nums[right] != tar {
				right--
			}
		}
		if nums[left] == tar && nums[right] == tar {
			lb, rb = left, right
			break
		}
	}
	return []int{lb, rb}
}
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target { //当相等的时候，左右扩展，直到边界值不等于tar
			leftBorder, rightBorder := mid-1, mid+1
			for leftBorder >= 0 && nums[leftBorder] == target {
				leftBorder -= 1
			}
			for rightBorder < len(nums) && nums[rightBorder] == target {
				rightBorder += 1
			}
			return []int{leftBorder + 1, rightBorder - 1}

		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return []int{-1, -1}
}

/*
平方根
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
*/
func q3() {
	input := 2
	output := methodSquare(input)
	fmt.Println(output)
}
func methodSquare(x int) int {
	res := int(math.Pow(float64(x), 0.5))
	res = twoToFind(x)
	return res
}
func twoToFind(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid <= x {
			ans = mid // ans为最大的mid*mid<=x的值
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans

}

/*
 */
func q4() {
	input := 121
	output := methodIntSquare(input)
	fmt.Println(output)
}
func methodIntSquare(x int) bool {
	res := Find2(x)
	return res
}
func Find2(x int) bool {
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == x {
			return true
		}
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
