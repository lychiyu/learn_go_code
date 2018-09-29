package main

import "fmt"

/**
slice 是个值类型
slice本身没有数据，是对底层array的一个视图view
slice是可以向后扩展，不可以向前扩展
s[i]不可以超越len(s),向后扩展不可以超过底层数组cap(s)

slice添加元素时，如果超越了cap，系统会重新分配更大的底层数组
由于是值传递，所以必须接收append返回的值
s = append(s,val)


*/

// []int slice
func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])

	s1 := arr[2:]
	s2 := arr[:]
	updateSlice(s1)
	fmt.Println(s1, arr)
	updateSlice(s2)
	fmt.Println(s2, arr)

	fmt.Println("slice扩展")
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s3 := arr1[2:6] // [2 3 4 5]
	s4 := s3[3:5]   // [5 6]
	fmt.Println(
		len(s3),
		cap(s3),
	)
	fmt.Println(s3, s4)

	fmt.Println("slice高级操作")
	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	// s6、s7不再是对arr1的view，而是对一个新的array的view
	fmt.Println(s5, s6, s7)
	fmt.Println(arr1)
}
