package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n",
		s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")
	//var s []int //Zero value for slice is nil
	s := []int{}

	//100个奇数
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println("前100个奇数是")
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)
	//len(s3) = 10, cap(s3) = 32
	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("Copying slice")
	//copy s1 to s2
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	//...代表把这个数组里的值拆开，
	//func(slice []Type, elems ...Type) []Type
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	//去数据，按下标
	front := s2[0]
	//截取？ 新增一个view
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	//当切片的元素减少时，len减少，cap不会，增加都会增加
	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
