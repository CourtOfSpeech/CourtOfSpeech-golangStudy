package main

import (
	"bufio"
	"fmt"
	"hello/functional/fib"
	"os"
)

func tryDefer() {
	//defer 先进后出
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	fmt.Println(3)
	//panic("error occurred")
	fmt.Println(4)

	//defer 参数在执行语句时计算
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("print end")
		}
	}

}

func writeFile(fileName string) {
	//创建一个file
	//file, err := os.Create(fileName)
	//制造一个错误，用来测试错误处理
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)

	//自己new一个error
	//err = errors.New("this is error")

	if err != nil {
		//这就是一个处理
		//fmt.Println("file already exists")
		//更好一点的处理
		//fmt.Println("Error: ", err.Error())
		//获得错误的类型  OpenFile的错误类型是一个*os.PathError ，判断这个错误是否是，是就输出，不是就报错
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			//fmt.Println(pathError.Op, pathError.Path, pathError.Err)
			fmt.Printf("pathError.Op: %s, pathError.Path: %s, pathError.Err: %s\n",
				pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	//关闭资源
	defer file.Close()

	//写数据，用bufio速度更快
	writer := bufio.NewWriter(file) //这里只是写入了内存里面
	//写入文件
	defer writer.Flush()

	f := fib.Fibonacci()
	//写入前20个
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
