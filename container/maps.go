package main

import "fmt"

/**
map使用哈希表，key必须可以比较相等
除了slice、map、function的内建类型都可以作为key
Struct类型不包含上述类型，也可以作为key
*/
func main() {
	m := map[string]string{
		"name": "lychiyu",
		"age":  "20",
		"tech": "golang",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	fmt.Println(m2) // empty map
	fmt.Println(m2 == nil)

	var m3 map[string]int
	fmt.Println(m3) // nil
	fmt.Println(m3 == nil)

	fmt.Println("遍历map (map是无序的，不保证遍历顺序)")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("get value")
	fmt.Println(m["1name"] == "") // map获取不存在的key时是空字符串

	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("not exist")
	}

	fmt.Println("delete value from map")
	delete(m, "age")
	fmt.Println(m["age"] == "")
}
