package tree

import (
	"fmt"
)

type TreeNode struct {
	Data       any
	LeftChild  *TreeNode
	RightChild *TreeNode
}

func CreateBinaryTree(inputList []any) (node *TreeNode, newInput []any) {
	if len(inputList) == 0 {
		return
	}
	data := inputList[0]
	fmt.Println("data:", data)
	fmt.Println("inputList before:", inputList)
	newInput = inputList[1:]
	fmt.Println("inputList after:", newInput)
	if data == nil {
		return
	}
	node = &TreeNode{}
	node.Data = data
	node.LeftChild, newInput = CreateBinaryTree(newInput)
	fmt.Println("after left:", inputList)
	node.RightChild, newInput = CreateBinaryTree(newInput)
	fmt.Println("after right:", inputList)
	return
}

// PreOrder 前序遍历
func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(" ", node.Data)
	PreOrder(node.LeftChild)
	PreOrder(node.RightChild)
}

// 中序遍历
func InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	InOrder(node.LeftChild)
	fmt.Print(" ", node.Data)
	InOrder(node.RightChild)
}

// 后序遍历
func PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PostOrder(node.LeftChild)
	PostOrder(node.RightChild)
	fmt.Print(" ", node.Data)
}
