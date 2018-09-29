package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 省略初始条件
func convertToBin(n int) string {
	/**
	整数转为2进制字符串
	*/
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 省略初始条件和递增条件
func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	for {
		fmt.Println("forever")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0),
	)

	readFile("abc.txt")

	forever()
}
