package util

import (
	"strconv"
	"strings"
)

const (
	//how many character space for a subtree
	minTreeWidth = 4
	//how many space for a node value
	minNodeWidth = 3
	vert     = "_"
	horiz    = "|"
	//the height between levels
	height = 2
)

func PrintTree(node *TreeNode) {

}

type TreeWidthInfo struct {
	LeftWidth  int
	RightWidth int
	NodeWidth  int
	TotalWidth int
}

func calTreeWidth(root *TreeNode, stats map[*TreeNode]*TreeWidthInfo) *TreeWidthInfo {
	widthInfo := &TreeWidthInfo{
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

// print tree with padding inherited from parent nodes
func printNodeWithPadding(root *TreeNode, stats map[*TreeNode]*TreeWidthInfo, padding int) {

}

func sprintNode(v int64) string {
	slen := GetLenOfInt64(v)
	str := strconv.FormatInt(v, 10)
	if slen >= minNodeWidth {
		return str
	}
	left := (minNodeWidth - slen) / 2
	right := minNodeWidth - slen - left
	return strings.Repeat(vert, left) + str + strings.Repeat(vert, right)
}
