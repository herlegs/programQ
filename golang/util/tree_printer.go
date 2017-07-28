package util

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	//how many character space for a subtree
	minTreeWidth = 4
	//how many space for a node value
	minNodeWidth = 3
	vert         = "|"
	horiz        = "-"
	//connection of horizontal and vertical tree arm
	armCorner = "."
	//place holder, background for printed tree
	S = " "
)

func PrintTree(root *TreeNode) {
	stats := make(map[*TreeNode]*treeWidthInfo)
	calTreeWidth(root, stats)
	printTree(root, stats)
}

type treeWidthInfo struct {
	LeftWidth     int
	RightWidth    int
	NodeWidth     int
	TotalWidth    int
	PaddingLeft   int
	IsPlaceHolder bool
}

func calTreeWidth(root *TreeNode, stats map[*TreeNode]*treeWidthInfo) *treeWidthInfo {
	widthInfo := &treeWidthInfo{
		TotalWidth: minTreeWidth,
	}
	if root == nil {
		return widthInfo
	}
	widthInfo.NodeWidth = Max(int64(GetLenOfInt64(root.Val)), minNodeWidth)

	leftInfo := calTreeWidth(root.Left, stats)
	rightInfo := calTreeWidth(root.Right, stats)

	widthInfo.LeftWidth = leftInfo.TotalWidth
	widthInfo.RightWidth = rightInfo.TotalWidth
	widthInfo.TotalWidth = widthInfo.LeftWidth + widthInfo.NodeWidth + widthInfo.RightWidth

	stats[root] = widthInfo
	return widthInfo
}

// print tree
func printTree(root *TreeNode, stats map[*TreeNode]*treeWidthInfo) {
	if root == nil {
		return
	}
	queue := NewQueue()
	queue.Push(root)
	for {
		levelCount := queue.Size()
		isLevelEmpty := true //all nodes have been processed
		currLevelStr, nextLevelArmStr := "", ""
		for i := 0; i < levelCount; i++ {
			item := queue.Pop()
			node := item.(*TreeNode)
			levelStr, armStr := sprintTreeNode(node, stats)
			if node == nil { //empty node
				queue.Push(node)
				continue
			}
			nodeInfo := stats[node]
			if !nodeInfo.IsPlaceHolder { // not placeholder (for node width) push children into queue
				isLevelEmpty = false
				nodeInfo.IsPlaceHolder = true
				queue.Push(node.Left)
				queue.Push(node)
				queue.Push(node.Right)
			} else {
				queue.Push(node)
			}
			currLevelStr += levelStr
			nextLevelArmStr += armStr
		}
		fmt.Println(currLevelStr)
		fmt.Println(nextLevelArmStr)
		if isLevelEmpty {
			break
		}
	}

}

//return current level string and next level vertical arm
func sprintTreeNode(node *TreeNode, stats map[*TreeNode]*treeWidthInfo) (string, string) {

	if node == nil {
		return strings.Repeat(S, minTreeWidth), strings.Repeat(S, minTreeWidth)
	}
	nodeInfo := stats[node]
	if nodeInfo.IsPlaceHolder {
		return strings.Repeat(S, nodeInfo.NodeWidth), strings.Repeat(S, nodeInfo.NodeWidth)
	}
	str := ""
	armStr := ""
	if node.Left == nil {
		str += strings.Repeat(S, minTreeWidth)
		armStr += strings.Repeat(S, minTreeWidth)
	} else {
		armLeftPad := stats[node.Left].LeftWidth + stats[node.Left].NodeWidth/2
		str += strings.Repeat(S, armLeftPad) + armCorner + strings.Repeat(horiz, stats[node.Left].TotalWidth-armLeftPad-len(armCorner))
		armStr += strings.Repeat(S, armLeftPad) + vert + strings.Repeat(S, stats[node.Left].TotalWidth-armLeftPad-len(vert))
	}
	nodeValStr := sprintNodeValue(node.Val)
	str += nodeValStr
	armStr += strings.Repeat(S, len(nodeValStr))
	if node.Right == nil {
		str += strings.Repeat(S, minTreeWidth)
		armStr += strings.Repeat(S, minTreeWidth)
	} else {
		armRightPad := stats[node.Right].RightWidth + stats[node.Right].NodeWidth/2
		str += strings.Repeat(horiz, stats[node.Right].TotalWidth-armRightPad-len(armCorner)) + armCorner + strings.Repeat(S, armRightPad)
		armStr += strings.Repeat(S, stats[node.Right].TotalWidth-armRightPad-len(vert)) + vert + strings.Repeat(S, armRightPad)
	}
	return str, armStr
}

func sprintNodeValue(v int64) string {
	slen := GetLenOfInt64(v)
	str := strconv.FormatInt(v, 10)
	if slen >= minNodeWidth {
		return str
	}
	left := (minNodeWidth - slen) / 2
	right := minNodeWidth - slen - left
	return strings.Repeat(S, left) + str + strings.Repeat(S, right)
}
