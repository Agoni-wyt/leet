package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	res := section(nums1, nums2)
	fmt.Println(res)
}

/*
题意：给定两个数组，编写一个函数来计算它们的交集。
输出结果中的每个元素一定是唯一的。 我们可以不考虑输出结果的顺序
例：
1.
输入：nums1 =[1,2,2,1],nums2=[2,2]
输出：2
2.
输入：nums1 =[4,9,5],nums2= [9,4,9,8,4]
输出：[9,4]
*/
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}) // 用map模拟set
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
func section(nums1, nums2 []int) []int {
	minList, toMapList := func(nums1, nums2 []int) ([]int, []int) {
		if len(nums1) > len(nums2) {
			return nums2, nums1
		} else {
			return nums1, nums2
		}
	}(nums1, nums2)
	set := make(map[int]bool)
	res := make([]int, 0)
	for _, v := range toMapList {
		if _, ok := set[v]; !ok {
			set[v] = true
		}
	}

	for _, v := range minList {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res

}
