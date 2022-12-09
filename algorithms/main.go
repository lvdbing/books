package main

import (
	"fmt"

	"github.com/lvdbing/books/algorithms/tree"
)

func main() {
	inputList := []any{3, 2, 9, nil, nil, 10, nil, nil, 8, nil, 4}
	// inputList := []any{3, 2, 9}
	treeNode, _ := tree.CreateBinaryTree(inputList)
	fmt.Println(*treeNode)
	fmt.Println("前序遍历：")
	tree.PreOrder(treeNode)
	fmt.Println()
	fmt.Println("中序遍历：")
	tree.InOrder(treeNode)
	fmt.Println()
	fmt.Println("后序遍历：")
	tree.PostOrder(treeNode)
}
