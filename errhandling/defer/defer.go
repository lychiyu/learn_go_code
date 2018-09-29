package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
defer 在函数return前执行，按照栈来执行, 参数在defer语句计算
 */

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 100; i++ {
		fmt.Fprintln(writer, i)
	}

}

func main() {
	tryDefer()

	writeFile("defer.txt")
}
