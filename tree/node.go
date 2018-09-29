package main

import "fmt"

/**
go语言中的面向对象只有封装没有继承和多态，所以使用struct
不论是地址还是结构体本身一律使用.来进行访问
 */

type treeNode struct {
	value       int
	left, right *treeNode
}

// struct没有构造函数，因为本身就提供了多种方法来进行实例化，可以使用外部的工厂函数来进行构造
func createNode(value int) *treeNode  {
	return &treeNode{value:value}
}
// 结构体的方法定义在外部
func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) traverse() {
	if node == nil{
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	root := treeNode{value: 3}
	root.left = &treeNode{4, nil, nil}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.right.right=createNode(2)
	fmt.Println(root.value, root.left, root.right)

	nodes := []treeNode{
		{3,nil,nil},
		{},
		{4, &root, &root},
	}
	fmt.Println(nodes)

	// 结构体方法
	root.print()

	root.traverse()
}
