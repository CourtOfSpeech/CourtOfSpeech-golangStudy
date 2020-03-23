package main

import "fmt"

type student struct {
	name  string
	phone int
	addr  string
}

func main() {
	fmt.Println("hello world!")

	student := student{}
	student.name = "jiangxin"
	student.phone = 18328753072
	student.addr = "四川绵阳三台县"
	fmt.Println(student.name)
	fmt.Println(student.phone)
<<<<<<< HEAD
	fmt.Println(student.addr)
=======
>>>>>>> fead5424d383b9d340c7f70b5214a963e9840865
	fmt.Println(student)

}
