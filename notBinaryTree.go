package main

import (
	"fmt"
)

type nodeTree struct {
	Node     int
	Children []*nodeTree //slice of pointers to nodeTree objects
}

func insert(root *nodeTree, parent int, value int) {
	if root == nil {
		return
	}
	if root.Node == parent {
		root.Children = append(root.Children, &nodeTree{Node: value, Children: []*nodeTree{}})
	} else {
		for _, child := range root.Children {
			insert(child, parent, value)
		}
	}
}

func printTree(root *nodeTree, level int) {
	if root == nil {
		return
	}
	fmt.Println(root.Node)
	for _, child := range root.Children {
		printTree(child, level+1)
	}
}

func height(root *nodeTree) int {
	if root == nil {
		return 0
	}

	maxChildHeight := 0
	for _, child := range root.Children {
		childHeight := height(child)
		if childHeight > maxChildHeight {
			maxChildHeight = childHeight
		}
	}

	return 1 + maxChildHeight
}

func main() {
	root := &nodeTree{Node: 1, Children: []*nodeTree{}}
	insert(root, 1, 3)
	insert(root, 1, 5)
	insert(root, 3, 6)
	insert(root, 3, 11)
	insert(root, 3, 13)

	fmt.Println("Non-Binary Tree with Slices:")
	printTree(root, 0)

	treeHeight := height(root)
	fmt.Println("Height of the tree:", treeHeight)
}
