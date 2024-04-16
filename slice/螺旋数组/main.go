package main

import "fmt"

func main() {

}

/*
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/

func q1() {
	input := 3
	output := generateMatrix(input)
	fmt.Println(output)
}
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	top, left, right, bottom := 0, 0, n-1, n-1
	num := 1
	tar := n * n
	for num <= tar {
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			ans[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i++ {
			ans[i][left] = num
			num++
		}
		left++
	}
	return ans
}
