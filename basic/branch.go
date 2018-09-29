package main

import (
	"fmt"
	"io/ioutil"
)

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("不支持的操作符: " + op)
	}
	// switch会自己break 除非使用fallthrough
	// switch 后面可以没有表达式，可以在case后
	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score <= 60:
		g = "F"
	case score < 80:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("错误的分数: %d", score))
	}
	return g
}

func main() {
	const filename = "./abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	// 另一种方法
	/*
		if 语句中是可以赋值的
		if 语句中赋值的变量作用域只在这个if语句中
	*/
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	eval(3, 2, "+")

	fmt.Println(grade(79))

	sum := 0
	/*
		for的条件里不需要括号
		for的条件里可以省略初始条件，结束条件，递增表达式
	*/
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
