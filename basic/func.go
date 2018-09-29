package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func evalOp(a, b int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// _ 不接收
		result, _ = div(a, b)
		return result, nil
	default:
		return 0, fmt.Errorf("不支持的操作符: %s", op)
	}
}

// 函数返回多个值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("function %s with args (%d, %d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	// 可变参数列表
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	rlt, err := evalOp(3, 1, "/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rlt)

	fmt.Println(
		div(10, 3),
	)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
