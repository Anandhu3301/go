package utils

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) CreateNode(new_data int) {
	newNode := &Node{Value: new_data}

	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (list *LinkedList) AddNodeIndex(afterval int, new_val int) {
	current := list.Head
	for current != nil && current.Value != afterval {
		current = current.Next
	}

	if current == nil {
		fmt.Println("Your Entered node value",afterval,"cannot be found")
		return
	}
	new_node := &Node{Value: new_val}

	new_node.Next = current.Next
	current.Next = new_node

}

func (list *LinkedList) AddNode(val int) {
	current := list.Head
	if current == nil {
		list.Head = &Node{Value: val}
	}
	for current.Next != nil {
		current = current.Next
	}
	new_node := &Node{Value: val}
	current.Next = new_node
}

func (list *LinkedList) Delete(value int) {
	if list.Head == nil {
		fmt.Println("The list is empty")
		return
	}

	if list.Head.Value == value {
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	var previous *Node
	for current.Next != nil && current.Value != value {
		previous = current
		current = current.Next
	}

	if current == nil {
		fmt.Println("Couldn't find the node")
		return
	}
	previous.Next = current.Next

}





// package utils

// import "fmt"

// type Node struct {
// 	Value int
// 	Next *Node
// }

// type LinkedList struct {
// 	Head *Node
// }

// func (list *LinkedList) Initialise(val int) {
// 	new_node := &Node{Value: val}

// 	if list.Head == nil {
// 		list.Head = new_node
// 	} else {
// 		current := list.Head
// 		for current.Next != nil {
// 			current = current.Next
// 		}
// 		current.Next = new_node
// 	}
// }

// func (list *LinkedList) AddAtIndex(searchval int , value int) {
// 	current := list.Head
// 	for current != nil && current.Value != searchval {
// 		current = current.Next
// 	}
// 	if current.Next == nil {
// 		fmt.Println("The entered search value cannot be found")
// 		return
// 	}
// 	new_node := &Node{Value: value}

// 	new_node.Next = current.Next
// 	current.Next = new_node
// }

