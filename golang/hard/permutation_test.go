package hard

import (
	"testing"
	"github.com/stretchr/testify/require"
	"fmt"
	"github.com/herlegs/programQ/golang/util"
	"regexp"
)

func TestSubset(t *testing.T) {
	dataSet := []struct {
		str string
		len int
		expected []string
	}{
		{"abc",1, []string{"a", "b", "c"}},
		{"abcd",2, []string{"ab", "bc", "cd", "ac", "ad", "bd"}},
		{"abb",2, []string{"ab", "bb"}},
		{"bbc",2, []string{"bb", "bc"}},
		{"abbc",2, []string{"ab", "ac","bb","bc"}},
		{"abbc",3, []string{"abb", "abc","bbc"}},
	}

	for _, data := range dataSet {
		out := Subset(data.str, data.len)
		fmt.Println(out)
		require.True(t, util.StringListEqual(out, data.expected), "source: %v, out: %v", data.str, out)
	}
}

func TestPermutation(t *testing.T) {
	dataSet := []struct {
		str string
		expected []string
	}{
		{"abc", []string{"abc", "acb", "bac", "bca", "cab","cba"}},
		{"abb", []string{"abb", "bab","bba"}},
		{"abba", []string{"aabb", "abab","abba","baab","baba","bbaa"}},
	}

	for _, data := range dataSet {
		out := Permutation(data.str)
		fmt.Println(out)
		require.True(t, util.StringListEqual(out, data.expected), "source: %v, out: %v", data.str, out)
	}
}

func TestSS(t *testing.T) {
	s := "Mr. Leonard Spock"
	re1, _ := regexp.Compile(`(Mr)(s)?\. (\w+) (\w+)`)
	result:= re1.FindStringSubmatch(s)
	fmt.Println(len(result))

	for k, v := range result {
		fmt.Printf("%d. %s\n", k, v)
	}
}

