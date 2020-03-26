package main

import (
	"fmt"
	"sync"
	"time"
)

type atominInt struct {
	value int
	//sync.Mutex: 互斥锁
	lock sync.Mutex
}

// *atominInt 表示这里是引用，也就是指针，而不是值传递，是传的一个内存地址
func (a *atominInt) increment() {
	fmt.Println("safe increment")
	//匿名函数
	func() {
		a.lock.Lock()
		//defer 声明的代码会在改函数结束之前的时间执行，而且不论函数是否出错都会执行
		//而且defer 最好放在前面，函数报错之前才会知道要执行这这行代码，或者会执行不到
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atominInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atominInt
	a.increment()
	//go 关键字声明代表在当前函数进程之外新开一个进程，并发任务时常用
	//且主函数停止后，该进程也会停止
	go func() {
		defer fmt.Println(`go`)
		a.increment()
	}()
	//睡眠1毫秒
	time.Sleep(time.Millisecond)
	fmt.Println(a.get)
}
