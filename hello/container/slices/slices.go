package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func updateArray(arr [8]int) {
	arr[1] = 111
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	//区间[),包含左，不包含右
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Printf("arr[:6] = %v\n", arr[:6])

	//切片即数组的一个view
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)

	s2 := arr[:]
	//n 为返回的字节数
	n, err := fmt.Println("s2 = ", s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	//golang中的切片slice底层通过数组实现，slice类似一个结构体，其中一个字段保存的是底层数组的地址，还有长度(len) 和 容量(cap）两个字段。
	//结构体作为函数参数时是值拷贝，同理，实际上slice作为函数参数时也是值拷贝，在函数中对slice的修改是通过slice中保存的地址对底层数组进行修改，所以函数外的silce看起来被改变了。
	//当需要对slice做插入和删除时，由于需要更改长度字段，值拷贝就不行了，需要传slice本身在内存中的地址。
	fmt.Println("after updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	//将底层数组作为参数时，是不能改变函数外的值，只有传入地址才可以改变
	fmt.Println("after updateArray()")
	updateArray(arr)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	//将数组还原
	arr[0], arr[2] = 0, 2
	fmt.Println("arr = ", arr)
	s1 = arr[2:6]
	s2 = s1[3:5]
	//因为切片只是底层数组的一个view，所以这里s2 = s1[3:6] ,不会报错
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	//当切片cap不够时，继续添加会新开一块内存，放这个新的底层数组
	fmt.Printf("s3 = %v\ns4 = %v\ns5 = %v\n",
		s3, s4, s5)
	//s4 and s5 并不是arr这个数组的 view  而是一个新的数组的 view
	fmt.Println("arr = ", arr)

}
