package main

import (
	"fmt"
	"os"
	"strconv"
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
func printFile(fileName string)  {
	os.Open(fileName)
	
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
}


