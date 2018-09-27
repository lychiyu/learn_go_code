package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "yes我爱学习!"     // utf8编码
	fmt.Println(len(s)) // 字节长度
	fmt.Printf("%s\n", []byte(s))
	fmt.Printf("%X\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()
	// 获取字符数量
	fmt.Println(utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()
	/**
	strings.Fields()
	strings.Split()
	strings.Join()

	strings.Contains()
	strings.Index()

	strings.ToLower()
	strings.ToUpper()

	strings.Trim()
	strings.TrimRight()
	strings.TrimLeft()
	 */

}
