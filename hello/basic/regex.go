package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse111@gmail.com"
const text1 = `My email is ccmouse111@gmail.com
email1 is abc@def.org
email2 is     kkk@qq.com
email3 is 	ddd@ab.com.cn
`

func main() {
	//func regexp.Compile(expr string) (*regexp.Regexp, error)
	//*regexp.Regexp 正则表达式的匹配器
	//re, err := regexp.Compile("ccmouse@gmail.com")
	//这种方式，如果传入的东西不符合正则表达式，会直接panic
	//re := regexp.MustCompile(".+@.+\\..+")
	//这种方式不用取转义
	//re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	//这里是根据上面的正则表达式在传入的字符串里找，并返回,这里只会返回一个
	//match := re.FindString(text)
	//FindAllString 可以找多个，-1代表找到全部匹配的
	///match := re.FindAllString(text1, -1)
	//fmt.Println(match)

	//能提取出括号中的内容
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text1, -1)
	fmt.Println(match)

	for _, m := range match {
		fmt.Println(m)
	}
}
