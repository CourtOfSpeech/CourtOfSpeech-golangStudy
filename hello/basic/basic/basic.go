package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//变量的定义
var (
	str   string
	index int
	flage bool
)

//变量出始值
func variableZeroVaue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

//变量定义时赋值
func varialeInitialValue() {
	var a, b int = 5, 6
	var s string = "xinjiang"
	fmt.Println(a, b, s)
}

//自动推导变量的类型
func varibleTypeDeduction() {
	var a, b, c, s = 3, 4.44, false, "esc"
	fmt.Println(a, b, c, s)
}

//自动推导变量的类型 shorter
func variableShorter() {
	a, b, c, s := 3, 4.44, true, "update"
	fmt.Println(a, b, c, s)
}

//欧拉方程 复数
func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

//勾股定理
func triangle() {
	a, b := 3, 4
	c := calcTriangle(a, b)
	fmt.Println(c)
}
func calcTriangle(a, b int) (c int) {
	c = int(math.Sqrt(float64(a*a + b*b)))
	return
}

//常量

func consts() {
	const (
		fileName = "abc.txt"
		a, b     = 3, 4
	)
	c := math.Sqrt(a*a + b*b)
	fmt.Println(fileName, c)
}

//枚举
func enums() {
	const (
		cpp = iota
		_
		golang
		java
		javascript
		c
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, golang, java, javascript, c)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

//函数主入口
func main() {
	variableZeroVaue()
	varialeInitialValue()
	varibleTypeDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	enums()

	if _, err := testTringle(); err != nil {
		fmt.Println(err)
	}
}

