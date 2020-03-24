package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//switch选择的用法
func eval(a, b int, op string) (int, error) {
	switch {
	case op == "+":
		return a + b, nil
	case op == "-":
		return a - b, nil
	case op == "*":
		return a * b, nil
	case op == "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)

	}
}

//除 和 余
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//不定参数
func sum(numbers ...int) (s int) {
	for _, number := range numbers {
		s += number
	}
	return
}

//值传递： 交换2个数的值
func swap(a, b int) (int, int) {
	return b, a
}

//函数做为参数
func apply(op func(int, int) int, a, b int) int {
	//通过反射来获取传入函数的对象
	p := reflect.ValueOf(op).Pointer()
	//通过对象获取这个函数的名称
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func main() {
	if result, err := eval(5, 5, "/"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(18, 5)
	fmt.Printf("q = %d, r = %d\n", q, r)

	fmt.Printf("1+2+...+5 = %d\n", sum(1, 2, 3, 4, 5))

	a, b := 19, 37
	a, b = swap(a, b)
	fmt.Printf("a, b 交换后的值为: a = %d, b = %d\n", a, b)

	res := apply(
		func(a, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4)

	fmt.Println("pow(3, 4) is:", res)

}
