package main

import "fmt"

var lastOccurred = make([]int, 0xffff)

//返回字符串中最长的不重复字符的长度 leetcode题目：寻找最长不含有重复字符的字串
func lenthOfNonRepeatingSubStr(s string) (maxLength int) {
	//对于每一个字符x
	//1.lastOccurred[x]不存在，或者 < start  =>无需操作
	//2.lastOccurred[x] >= start	=>更新start
	//以上不论那种情况
	//更新lastOccurred[x]，更新maxLength
	//lastOccurred := make(map[rune]int)
	//根据pprof性能调优 提示 map 用的时间比较久，这里换一种方式
	//由于int的默认都为0  所以Occurred+1
	//lastOccurred := make([]int, 0xffff) //换成数组，长度是0xffff, 但是这里每次执行方法都会产生新的数组，go的垃圾回收就会占很多时间，所以把它定义在函数外面

	//如果将数组定义在外面，则需要每次都将上次的数据清空

	for i := range lastOccurred {
		lastOccurred[i] = 0
	}

	start := 0 //当前找到不含重复字串的开始位置

	//rune汉字也就只有一个字符，不然会出错
	for i, ch := range []rune(s) {
		//if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
		if lastI := lastOccurred[ch]; lastI > start {
			start = lastI

		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		//lastOccurred[ch] = i
		lastOccurred[ch] = i + 1

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
