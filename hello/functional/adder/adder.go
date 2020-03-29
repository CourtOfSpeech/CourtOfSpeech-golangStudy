package main

import "fmt"

//adder()返回一个有返回值的函数，这个函数用来累加,这个函数叫一个闭包
//函数体：1.局部变量 v， 2.自由变量 sum
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//返回值 int 即当前累加完的值
//iAdder 即下一个返回函数
//这里是递归
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		//fmt.Println(a(i))
		fmt.Printf("0+...+ %d = %d\n", i, a(i))
	}
	fmt.Println("正统一点的函数编程")
	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0+...+ %d = %d\n", i, s)
	}
}
