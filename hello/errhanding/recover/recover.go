package main

import (
	"fmt"
)

//一个panic
func tryRecover() {
	//recover 只能在defer中使用
	defer func() {
		//r := recover()
		//如果程序没有panic ，就不做处理
		if r := recover(); r != nil {
			//如果r 是一个error，则处理它
			if err, ok := r.(error); ok {
				fmt.Println("Error occurred : ", err)
			} else {
				//否则继续panic
				//panic(r)
				panic(fmt.Sprintf("I don't know what to do : %v", r))
			}
		}
	}()

	//panic(errors.New("this is an error"))

	// b := 0
	// a := 5 / b
	// fmt.Println(a)

	//由于这里不是一个error ，所以程序会挂掉
	panic(123)
}

//recover 的使用
func main() {
	tryRecover()
}
