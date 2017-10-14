package util

import (
	"testing"
)

func TestPrintTree(t *testing.T) {
	root := createSampleTree()
	PrintTree(root)
}


func createSampleTree() *TreeNode {
	root := createRandomTreeNode()
	root.Left = createRandomTreeNode()
	root.Right = createRandomTreeNode()

	root.Left.Right = createRandomTreeNode()
	//root.Left.Right = createRandomTreeNode()

	root.Right.Left = createRandomTreeNode()

	return root
}

func createRandomTreeNode() *TreeNode {
	return &TreeNode{
		Val: GetRandomInt64(1, 10000),
	}
}
