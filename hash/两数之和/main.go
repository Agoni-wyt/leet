package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	output := twoSum(nums, target)
	fmt.Println(output)

}

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for i, num := range nums {
		if v, ok := set[target-num]; ok {
			return []int{i, v}
		} else {
			set[num] = i
		}
	}
	return nil
}
