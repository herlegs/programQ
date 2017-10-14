package util

import (
	"fmt"
	"strings"
	"strconv"
)

const(
	nullNode = "#"
	sep = ","
)

func Serialize(root *TreeNode) string {
	if root == nil {
		return nullNode
	}
	str := ""
	queue := NewQueue()
	queue.Push(root)
	//total nodes in queue needs to be processed
	nodeCount := 0
	if root != nil {
		nodeCount = 1
	}
	for !queue.IsEmpty() {
		node := queue.Pop().(*TreeNode)
		nodeStr := ""
		if node != nil {
			nodeCount--
			nodeStr = fmt.Sprintf("%v",node.Val)
			queue.Push(node.Left)
			queue.Push(node.Right)
			if node.Left != nil {
				nodeCount++
			}
			if node.Right != nil {
				nodeCount++
			}
		} else {
			nodeStr = nullNode
		}
		if str != "" {
			str += sep
		}
		str += nodeStr
		if nodeCount == 0 {
			break;
		}
	}
	return str
}

func Deserialize(s string) *TreeNode {
	elems := strings.Split(s, sep)
	queue := NewQueue()
	var root *TreeNode
	leftAdded := false
	for _, ele := range elems {
		cur := makeNode(ele)
		if root == nil {
			root = cur
		}
		if !queue.IsEmpty() {
			parent := queue.Peek().(*TreeNode)
			if !leftAdded {
				parent.Left = cur
				leftAdded = true
			} else {
				parent.Right = cur
				queue.Pop()
				leftAdded = false
			}
		}
		if cur != nil {
			queue.Push(cur)
		}
	}
	return root
}

func makeNode(val string) *TreeNode {
	if val == nullNode {
		return nil
	} else{
		val, _ := strconv.ParseInt(val, 10, 64)
		return &TreeNode{
			Val: val,
		}
	}
}