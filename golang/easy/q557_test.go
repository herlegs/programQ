package easy

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestReverseWords(t *testing.T) {
	input := "Let's take LeetCode contest"
	expected := "s'teL ekat edoCteeL tsetnoc"
	require.Equal(t,expected, ReverseWords(input))

	input = ""
	expected = ""
	require.Equal(t,expected, ReverseWords(input))
}
