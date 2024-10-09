package utils


type Node struct {
	Value int
	Next *Node
}

type LinkedList struct{
	Head *Node
}

func (list *LinkedList) CreateNode(new_data  int) {
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