package main

import "fmt"

/**
数组是值类型
在go语言中一般不直接使用数组，而是使用切片
*/

func printArray(arr [5]int) {
	for i, v := range arr {
		println(i, v)
	}
	arr[0] = 100
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}

	arr3 := [...]int{4, 5, 6, 7, 8}
	// 4行5列
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(arr3[:2])
	fmt.Println(grid)

	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		println(i, v)
	}

	// 会进行数组拷贝
	printArray(arr1)
	printArray(arr3)

	fmt.Println(arr1, arr3)
}
