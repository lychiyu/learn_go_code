package main

import "fmt"

/**
获取最长无重复的子串长度
 */
func lengthOfNoneRepeatSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start, maxLength := 0, 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
func main() {
	fmt.Println(lengthOfNoneRepeatSubStr("abcabca"))
	fmt.Println(lengthOfNoneRepeatSubStr("bbbb"))
	fmt.Println(lengthOfNoneRepeatSubStr(""))
	fmt.Println(lengthOfNoneRepeatSubStr("abcdefghi"))
	fmt.Println(lengthOfNoneRepeatSubStr("b"))
	fmt.Println(lengthOfNoneRepeatSubStr("yes我爱学习!"))
	fmt.Println(lengthOfNoneRepeatSubStr("黑化黑灰化肥黑灰会挥发发灰黑化肥黑灰化肥挥发灰化灰黑化肥灰黑会发挥发黑灰化肥灰黑化肥发挥!"))
}
