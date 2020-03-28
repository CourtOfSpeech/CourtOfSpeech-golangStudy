package main

import (
	"fmt"
	"hello/queue/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	//fmt.Println(q)
	q.Push("adsfs")
	fmt.Println(q.Pop())
	
}
