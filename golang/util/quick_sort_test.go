package util

import (
	"testing"
	"fmt"
	"runtime"
	"github.com/stretchr/testify/require"
)

var testList []int

func init(){
	fmt.Println(runtime.NumCPU())
	testList = []int{}
	for i := 1000000; i > 0; i-- {
		testList = append(testList, i)
	}
}

func TestQSort(t *testing.T) {
	list := []int{3,2,4,5,7,5,9,12}
	QSortM(list)
	require.EqualValues(t, []int{2,3,4,5,5,7,9,12}, list)
	fmt.Println(list)

	list = []int{3}
	QSortM(list)
	require.EqualValues(t, []int{3}, list)
	fmt.Println(list)

	list = []int{}
	for i := 10000; i > 0; i-- {
		list = append(list, i)
	}
	QSortM(list)
	fmt.Println(list)
}

/*
100	  12525902 ns/op
*/
func BenchmarkQSort(b *testing.B) {
	fmt.Println(len(testList))
	for i := 0; i < b.N; i++ {
		QSort(testList)
	}
}

/***
300	 114866395 ns/op
2000	   4088127 ns/op --- 1/8
 */
func BenchmarkQSortM(b *testing.B) {
	fmt.Println(len(testList))
	for i := 0; i < b.N; i++ {
		QSortM(testList)
	}
}
