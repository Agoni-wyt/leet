package main

import "fmt"

func main() {

	var str string
	var target int

	fmt.Println(&target)
	fmt.Println(&str)
	strByte := []byte(str)

	reverse(strByte, 0, len(strByte)-1)      //全部翻转
	reverse(strByte, 0, target-1)            // 翻转前区间 ，长度n
	reverse(strByte, target, len(strByte)-1) // 翻转后区间，长度length-n

	fmt.Printf(string(strByte))

}

/*
字符串的右旋转操作是把字符串尾部的若干个字符转移到字符串的前面。给定一个字符串 s 和一个正整数 k，请编写一个函数，将字符串中的后面 k 个字符移到字符串的前面，实现字符串的右旋转操作。

例如，对于输入字符串 "abcdefg" 和整数 2，函数应该将其转换为 "fgabcde"。

输入：输入共包含两行，第一行为一个正整数 k，代表右旋转的位数。第二行为字符串 s，代表需要旋转的字符串。

输出：输出共一行，为进行了右旋转操作后的字符串。

样例输入：

2
abcdefg
样例输出：

fgabcde
数据范围：1 <= k < 10000, 1 <= s.length < 10000
思路就是 通过 整体倒叙，把两段子串顺序颠倒，两个段子串里的的字符在倒叙一把，负负得正，这样就不影响子串里面字符的顺序了。
*/
func reverse(strByte []byte, l, r int) {
	for l < r {
		strByte[l], strByte[r] = strByte[r], strByte[l]
		l++
		r--
	}
}
