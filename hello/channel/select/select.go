package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)

			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %v\n", id, n)
	}
}

func createworker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	//var c1, c2 chan int //c1 and c2 ==nil
	c1, c2 := generator(), generator()
	worker := createworker(0)
	// n := 0
	// haseValue := false
	// for {
	// 	var activeWorker chan<- int
	// 	if haseValue {
	// 		activeWorker = worker
	// 	}
	// 	select {
	// 	case n = <-c1:
	// 		haseValue = true
	// 	case n = <-c2:
	// 		haseValue = true
	// 	case activeWorker <- n:
	// 		haseValue = false

	// 		// default:
	// 		// 	fmt.Println("Not Recevied Values")
	// 	}
	// }

	//更好的一种方法
	var values []int
	//计时器 time.After(时间) 过了多少时间，就会向	<-chan time.Time  这个channel中放数据
	tm := time.After(10 * time.Second)
	//定时
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeOut")
		case <-tick:
			fmt.Println("queue lne", len(values))
		case <-tm:
			fmt.Println("bay")
			return
		}
	}
}
