package main

import (
	"fmt"
	"sync"
)

// channel作为参数使用
func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		//通过channel通信来共享内存，这里告诉外面接收来消息
		go func() {
			//done <- true
			w.done()
		}()

	}
}

//用go 的 sync.WaitGroup 函数编程
type worker struct {
	in   chan int
	done func()
}

//channel作为返回值 这里返回的channel只能收数据
func createworker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	//新开一个sync.WaitGroup
	var wg sync.WaitGroup
	//多开几个协程来接收channel
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createworker(i, &wg)

	}
	//添加20个任务
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		//这里等打印完了以后才会结束函数，否则会阻塞在这里,但是这样会发一个in,打印一个
		//<-worker.done
		//wg.Add(1)
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-worker.done
	}

	//等待添加的任务结束
	wg.Wait()

	//先一次性把任务法出去，然这里在等待，全部打印结束后在退出,这里是我们自己写代码去等待，还可以用go的sync.WaitGroup去等待
	// for _, worker := range workers {
	// 	<-worker.done
	// 	<-worker.done
	// }
}

func main() {
	chanDemo()
}
