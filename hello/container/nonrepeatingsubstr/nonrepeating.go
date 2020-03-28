package main

import "fmt"

//返回字符串中最长的不重复字符的长度 leetcode题目：寻找最长不含有重复字符的字串
func lenthOfNonRepeatingSubStr(s string) (maxLength int) {
	//对于每一个字符x
	//1.lastOccurred[x]不存在，或者 < start  =>无需操作
	//2.lastOccurred[x] >= start	=>更新start
	//以上不论那种情况
	//更新lastOccurred[x]，更新maxLength
	lastOccurred := make(map[rune]int)
	start := 0 //当前找到不含重复字串的开始位置

	//rune汉字也就只有一个字符，不然会出错
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
	fmt.Println(
		lenthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lenthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(
		lenthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lenthOfNonRepeatingSubStr(""))
	fmt.Println(
		lenthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lenthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(
		lenthOfNonRepeatingSubStr(
			"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
