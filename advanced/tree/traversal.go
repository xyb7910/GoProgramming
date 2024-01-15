package tree

import "fmt"

func (node *Node) Traversal() {
	node.TraversalFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraversalFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraversalFunc(f)
	f(node)
	node.Right.TraversalFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraversalFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
