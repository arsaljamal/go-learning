package main

import (
	"fmt"
)

type TreeNode struct {
	Parent *TreeNode
}

// Function to find the depth (level) of a node in the tree.
func findDepth(node *TreeNode) int {
	depth := 0
	for node != nil {
		depth++
		node = node.Parent
	}
	return depth
}

// Function to find the common ancestor of two nodes in a tree.
func findCommonAncestor(node1, node2 *TreeNode) *TreeNode {
	// Find the depths of both nodes.
	depth1 := findDepth(node1)
	depth2 := findDepth(node2)
	firtPointer := node1
	secondPointer := node2
	// Adjust the positions of nodes to the same level.
	for depth1 > depth2 {
		firtPointer = firtPointer.Parent
		depth1--
	}
	for depth2 > depth1 {
		secondPointer = secondPointer.Parent
		depth2--
	}

	// Move both nodes up until they reach the common ancestor.
	for firtPointer != secondPointer {
		firtPointer = firtPointer.Parent
		secondPointer = secondPointer.Parent
	}

	return firtPointer // Common ancestor found
}

func main() {
	// Build a tree.
	root := &TreeNode{}
	node1 := &TreeNode{Parent: root}
	node2 := &TreeNode{Parent: root}
	node3 := &TreeNode{Parent: node1}
	node4 := &TreeNode{Parent: node2}
	node5 := &TreeNode{Parent: node3}

	// Example: Find common ancestor of node4 and node5
	commonAncestor := findCommonAncestor(node4, node5)

	if commonAncestor != nil {
		fmt.Printf("Common Ancestor: %+v\n", &commonAncestor)
	} else {
		fmt.Println("No common ancestor found.")
	}
}
