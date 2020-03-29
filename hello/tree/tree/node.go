package tree

import "fmt"

// Node
type Node struct {
	Value       int
	Left, Right *Node
}

//打印node节点
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//创建node节点
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//给node赋值
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}
