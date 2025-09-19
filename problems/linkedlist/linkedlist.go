package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}


func removeNthFromEnd(index int, node *Node) *Node {
	// Create a dummy node to handle cases where the head needs to be removed.
	head := &Node{Next: node}
	first := head
	second := head

	// Advance the first pointer by 'index+1' nodes.
	for i := 0; i <= index; i++ {
		if first == nil {
			return node // Index is out of range, no removal.
		}
		first = first.Next
	}

	// Move both pointers until the first pointer reaches the end.
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the node by updating the next pointer of the second-to-last node.
	second.Next = second.Next.Next

	return head.Next
	//Think of it like two runners 🏃‍♂️🏃‍♀️. 
	// You send one runner ahead by n steps. 
	// Then move both together until the fast one hits the end.
	// Now the slow one is exactly before the node to delete.
}


// Function to print the linked list.
func printLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	head.Next.Next.Next = &Node{Value: 4}
	head.Next.Next.Next.Next = &Node{Value: 5}

	fmt.Println("Original Linked List:")
	printLinkedList(head)

	indexToRemove := 3
	head = removeNthFromEnd(indexToRemove, head)

	fmt.Printf("Linked List after removing node at index %d from the end:\n", indexToRemove)
	printLinkedList(head)
}