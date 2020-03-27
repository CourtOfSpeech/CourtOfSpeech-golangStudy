package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	//数组的声明
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	//多维数组
	var grid [4][5]int

	fmt.Println("array definitions:")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printAarray(arr1)")
	printArray(arr1)

	fmt.Println("printAarray(arr3)")
	printArray(arr3)

	fmt.Printf("arr1 = %#v, arr3 = %#v\n", arr1, arr3)

}
