package main

import (
	"fmt"
	"time"
)

//运行 go run -race xxx.go  可以查看数据访问冲突
func main() {
	//var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i) //打印是io操作，会有goroutine的切换，会交出控制权
				//不交出控制权
				//a[i]++
				//主动交出控制权
				//runtime.Gosched()
			}
		}(i)
	}
	//毫秒
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
