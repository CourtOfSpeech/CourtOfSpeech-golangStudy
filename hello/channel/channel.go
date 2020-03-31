package main

import (
	"fmt"
	"time"
)

// //channel作为参数使用
func worker(id int, c chan int) {
	// for {
	// 	//n := <-c
	// 	//fmt.Println(n)
	// 	//判断channel是否还有值一种方式
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("worker %d received %v\n", id, n)
	// }

	//判断channel是否还有值另一种方式
	for n := range c {
		fmt.Printf("worker %d received %v\n", id, n)
	}
}

//channel作为返回值 这里返回的channel只能收数据
func createworker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %v\n", id, <-c)
		}
	}()
	return c
}

// func chanDemo() {
// 	// var c chan int	//c ==nil 这样就可以用了
// 	c := make(chan int)
// 	//channel 接收数据后，需要有其他goroutine去接收，不然会阻塞，报错
// 	go worker(0, c)
// 	//发数据到channel
// 	c <- 1
// 	c <- 2
// 	//睡眠一会，不然程序退出了，就只能打印一个数据
// 	time.Sleep(time.Microsecond)
// }

// func chanDemo2() {
// 	//多开几个协程来接收channel
// 	var channels [10]chan int
// 	for i := 0; i < 10; i++ {
// 		channels[i] = make(chan int)
// 		go worker(i, channels[i])
// 	}

// 	for i := 0; i < 10; i++ {
// 		channels[i] <- 'a' + i
// 	}
// 	time.Sleep(time.Millisecond)
// }

func chanDemo3() {
	//多开几个协程来接收channel
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createworker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

//有缓冲区的channel，当channel接收了一定数据才会阻塞，否则不会阻塞
func bufferedChannel() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	go worker(0, c)
	//超出channel的缓冲区 ：all goroutines are asleep - deadlock!
	//c <- 4
}

//关闭channel
func channelClose() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	go worker(0, c)
	//关闭channel，关闭后，channel读取方 读出的数据就是channel类型的ZeroValue,这里是 0
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//	chanDemo()
	//chanDemo2()
	//chanDemo3()
	//bufferedChannel()
	channelClose()
}
