package packet

import "fmt"

type Bst struct {
	Data  int
	Left  *Bst
	Right *Bst
}

type Root struct {
	Root *Bst
}

func GetnewNode(val int) *Bst {
	new_node := &Bst{Data: val, Left: nil, Right: nil}
	return new_node
}

func (r *Root) Insert( val int) *Bst {
	if r.Root == nil {
		r.Root = GetnewNode(val)
		return r.Root
	}
	return InsertNode(r.Root, val)

}

func InsertNode(node *Bst, val int) *Bst {
	if val < node.Data {
		if node.Left == nil {
			node.Left = GetnewNode(val)
		} else {
			InsertNode(node.Left, val)
		}
	} else if val > node.Data {
		if node.Right == nil {
			node.Right = GetnewNode(val)
		} else {
			InsertNode(node.Right, val)
		}
	}
	return node
}

func InorderTraversal(node *Bst) {
	if node != nil {
		InorderTraversal(node.Left)
		fmt.Print(node.Data, " ")
		InorderTraversal(node.Right)

	}
}