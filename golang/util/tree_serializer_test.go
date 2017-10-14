package util

import (
	"testing"
	"github.com/stretchr/testify/require"
	"fmt"
)

func TestSerializer(t *testing.T) {
	treeStr := "#"
	tree := Deserialize(treeStr)
	serializedStr := Serialize(tree)
	require.Equal(t, treeStr, serializedStr)
}

func TestSerialize(t *testing.T) {
	tree := createSampleTree()
	PrintTree(tree)
	str := Serialize(tree)
	fmt.Println(str)
}

func TestDeserialize(t *testing.T) {
	str := "1,2,3,#,4,5,#"
	tree := Deserialize(str)
	PrintTree(tree)
}