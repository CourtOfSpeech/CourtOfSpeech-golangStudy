package main

import (
	"bufio"
	"fmt"
	"hello/functional/fib"
	"io"
	"strings"
)

//为函数实现接口,只有是类型就可以实现接口
type intGen func() int

//函数实现接口
func (i intGen) Read(p []byte) (n int, err error) {
	next := i()
	//文件上限
	if next > 1000 {
		return 0, io.EOF
	}
	//具体实现太底层了，用Sprintf来实现一下
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

//扫描字符后输出
func printFileContents(reader io.Reader) {
	//bufio.NewScanner读取数据到缓冲区时，且想要采用分隔符分隔数据流时
	scanner := bufio.NewScanner(reader)
	//scanner.Scan() 迭代器
	for scanner.Scan() {
		//scanner.Text() 扫描文本，以行返回
		fmt.Println(scanner.Text())
	}
}

func main() {
	//f := fib.Fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
