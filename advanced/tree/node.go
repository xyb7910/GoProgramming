package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// (node Node)强调函数的使用者，为结构定义方法的接受者
func (node Node) Print() {
	fmt.Println(node.Value, " ")
}

func (node *Node) SetValue(value int) { //只有使用指针才可以改变结构内容
	if node == nil {
		fmt.Println("Setting Value to nil" + "node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node { //使用工厂函数，注意返回为局部变量的地址
	return &Node{Value: value}
}
