package util

import "container/list"

func TreeInorderRecursive(root *TreeNode) []int64 {
	res := []int64{}
	recursiveHelper(root, &res)
	return res
}

func recursiveHelper(root *TreeNode, res *[]int64){
	if root == nil {
		return
	}
	recursiveHelper(root.Left, res)
	*res = append(*res, root.Val)
	recursiveHelper(root.Right, res)
}

func TreeInorderIterative(root *TreeNode) []int64 {
	res := []int64{}
	if root == nil {
		return res
	}
	stack := NewStack()
	for {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		if stack.IsEmpty(){
			break
		}
		root = stack.Pop().(*TreeNode)
		res = append(res, root.Val)
		root = root.Right
	}
	return res
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



