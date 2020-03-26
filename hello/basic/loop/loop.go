package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//10进制转二进制
func convertToBin(n int) (result string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		//strconv.Itoa(lsb)将整数数字转化为字符串
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

//读取文件信息
func printFile(fileName string) {
	//os.open()打开文件，只能读取信息，不能修改
	if file, err := os.Open(fileName); err != nil {
		fmt.Println(err)
	} else {
		printFileContents(file)
	}

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
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(0),
	)
	// Uncomment to see it runs forever
	//forever()

	printFile("abc.txt")

	fmt.Println(`printing a string`)

	s := `abc"d"
	jiangx
	126%
	
	end`
	printFileContents(strings.NewReader(s))

}
