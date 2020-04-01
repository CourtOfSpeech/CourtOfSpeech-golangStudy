package main

import (
	"fmt"
	"hello/tree/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	//new(tree.Node)是一个空的
	root.Right.Left = new(tree.Node)
	//创建node并赋值
	root.Left.Right = tree.CreateNode(2)
	//给已经有的node赋值
	root.Right.Left.SetValue(4)

	fmt.Print("In-order traversal: ")
	root.Traverse()

	//数node节点的个数
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	//channel遍历二叉树,这里用来统计最大数
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}

	fmt.Println("Max Node value: ", maxNode)
}
