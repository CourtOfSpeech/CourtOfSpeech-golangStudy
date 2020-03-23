package main

import "fmt"

type student struct {
	name  string
	phone string
	addr  string
	age   int64
}

func main() {
	fmt.Println("hello world!")

	student := student{}
	student.name = "jiangxin"
	student.phone = "1832875xxxx"
	student.addr = "四川绵阳三台县"
	student.age = 18
	fmt.Println(student.name)
	fmt.Println(student.phone)
	fmt.Println(student.addr)
	fmt.Println(student.age)
	fmt.Println(student)
	

}
