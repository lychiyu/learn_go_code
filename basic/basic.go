package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func variableZeroValue() {
	// 未赋值变量
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	//初始化变量
	var a, b int = 3, 4
	var s string = "abc"
	// rune 字符类型
	var c rune = 'a'
	fmt.Printf("%d %d %q\n", a, b, s)
	fmt.Printf("%q %d\n", c, c)
}

func variableTypeDeduction() {
	// 自动推导变量类型
	var a, b, c, d = 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	// 函数内简洁声明变量方法 :=
	a, b, c, d := 3, 4, true, "abc"
	b = 5
	fmt.Println(a, b, c, d)
}

func euler()  {
	// 复数类型
	c := 3+4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi)+1)
}

// go语言常量定义
func consts()  {
	const filename string="test.txt"
	// 常量的数值可以作为任意类型使用 不需要强转
	const a, b= 1, 2
	c := a+b
	fmt.Println(c)
}

func enums()  {
	// 枚举类型
	const (
		//cpp = 0
		//java = 1
		//python = 2

		// iota表示示自增的种子默认是0
		cpp = iota
		java
		_	// 表示跳过
		python
		golang
	)

	const(
		b = 1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	fmt.Printf("hello world\n")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	euler()

	consts()

	enums()
}
