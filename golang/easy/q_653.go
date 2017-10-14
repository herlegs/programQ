package easy

import (
	"container/list"
)

/*
https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/

Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.
 */

type TreeNode struct {
	Val   int64
	Left  *TreeNode
	Right *TreeNode
}


func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	inorder := TreeInorderIterator(root)
	reverse := TreeReverseInorderIterator(root)
	s, e := inorder.Next(), reverse.Next()
	if s == e {
		return false
	}
	for inorder.HasNext() && reverse.HasNext() && s != e{
		sum := int(s.Val + e.Val)
		if sum == k {
			return true
		} else if sum < k {
			s = inorder.Next()
		} else {
			e = reverse.Next()
		}
	}
	return false
}

type TreeIterator interface {
	HasNext() bool
	Next() *TreeNode
}

func TreeInorderIterator(root *TreeNode) TreeIterator {
	return (&treeInorderIterator{}).init(root)
}

type treeInorderIterator struct {
	cur *TreeNode
	stack *list.List
}

func (i *treeInorderIterator) init(root *TreeNode) *treeInorderIterator {
	i.cur = root
	i.stack = list.New()
	i.prepare()
	return i
}

func (i *treeInorderIterator) HasNext() bool {
	return i.stack.Len() != 0
}

func (i *treeInorderIterator) prepare() {
	for i.cur != nil {
		i.stack.PushBack(i.cur)
		i.cur = i.cur.Left
	}
}

func (i *treeInorderIterator) Next() *TreeNode {
	nodeEle := i.stack.Back()
	if nodeEle != nil {
		i.stack.Remove(nodeEle)
		i.cur = nodeEle.Value.(*TreeNode)
		result := i.cur
		i.cur = i.cur.Right
		i.prepare()
		return result
	}
	return nil
}

func TreeReverseInorderIterator(root *TreeNode) TreeIterator {
	return (&treeReverseInorderIterator{}).init(root)
}

type treeReverseInorderIterator struct {
	cur *TreeNode
	stack *list.List
}

func (i *treeReverseInorderIterator) init(root *TreeNode) *treeReverseInorderIterator {
	i.cur = root
	i.stack = list.New()
	i.prepare()
	return i
}

func (i *treeReverseInorderIterator) HasNext() bool {
	return i.stack.Len() != 0
}

func (i *treeReverseInorderIterator) prepare() {
	for i.cur != nil {
		i.stack.PushBack(i.cur)
		i.cur = i.cur.Right
	}
}

func (i *treeReverseInorderIterator) Next() *TreeNode {
	nodeEle := i.stack.Back()
	if nodeEle != nil {
		i.stack.Remove(nodeEle)
		i.cur = nodeEle.Value.(*TreeNode)
		result := i.cur
		i.cur = i.cur.Left
		i.prepare()
		return result
	}
	return nil
}