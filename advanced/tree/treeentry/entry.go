package main

import (
	"LearingGo/advanced/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	//后续遍历：左 右 根
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	root := tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}

	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Println("In-order traversal: ")
	root.Traversal()

	fmt.Println("My own post-order traversal: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraversalFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNodeValue := 0
	for node := range c {
		if node.Value > maxNodeValue {
			maxNodeValue = node.Value
		}
		fmt.Println("Max node value:", maxNodeValue)
	}
}
