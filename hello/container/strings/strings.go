package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "go 使用 utf-8的编码方式"
	fmt.Println(str)

	for _, b := range []byte(str) {
		//%X 十六进制表示，字母形式为大写 A-F
		fmt.Printf("%X", b)
	}
	fmt.Println()

	// ch is a rune
	for i, ch := range str {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	//utf8.RuneCountInString() 返回字符串的长度
	fmt.Println("Rune count:", utf8.RuneCountInString(str))

	//转成一个一个的字符
	bytes := []byte(str)
	for len(bytes) > 0 {
		//返回b的第一个rune值和它所占用的字节数；
		//如果b不是以一个有效的rune值开头，
		//返回U+FFFD（字符?的Unicode）和0
		ch, size := utf8.DecodeRune(bytes)
		//从新定义视图，即新切片
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()
	
	for i, ch := range []rune(str) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}
