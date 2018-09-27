package main

import "fmt"

// slice的cap每次翻倍
func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))
}

func main() {

	fmt.Printf("创建slice \n")
	// 定义一个slice， 默认是nil
	var s [] int
	fmt.Println(s == nil) // true

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)


	s1 := []int{2, 4, 6, 8}
	printSlice(s1)	// len=4, cap=4

	s2 := make([]int, 16)
	printSlice(s2)	//len=16, cap=16


	s3 := make([]int, 10, 32)
	printSlice(s3)	// len=10, cap=32


	fmt.Printf("copy slice \n")
	copy(s2, s1)
	fmt.Println(s2)	// [2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0]
	printSlice(s2)	// len=16, cap=16

	fmt.Printf("delete element from slice \n")
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)	// [2 4 6 0 0 0 0 0 0 0 0 0 0 0 0]
	printSlice(s2)	// len=15, cap=16
	fmt.Printf("pop from front \n")
	front := s2[0]
	s2 = s2[1:]

	fmt.Printf("pop from back \n")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(front, tail)
	fmt.Println(s2)	// [4 6 0 0 0 0 0 0 0 0 0 0 0]
	printSlice(s2)	// len=13, cap=15

}
