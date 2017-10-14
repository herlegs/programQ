package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"strconv"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

func TestBinarySearch(t *testing.T) {
	require := require.New(t)
	var input []int
	var target int
	var result int

	input = []int{1, 4, 6, 12, 30, 68}
	target = 5

	result = BinarySearchGreaterThan(input, 0, len(input)-1, target)
	require.Equal(2, result)

	result = BinarySearchLessThan(input, 0, len(input)-1, target)
	require.Equal(1, result)

	result = BinarySearchGreaterThan(input, 2, len(input)-1, target)
	require.Equal(2, result)

	result = BinarySearchLessThan(input, 2, len(input)-1, target)
	require.Equal(1, result)
}

func TestBSFirstBigger(t *testing.T) {
	//need to find last element equal to target or bigger than target
	//eg. for questions: how many elements are there equal or less than target?
	//in this case return the first found equal element index does not work (might be duplicate)
	//return the last equal or first bigger will do
	input := []int{1, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 12, 30, 68}
	target := 5
	result := BSLastSmaller(input, 0, len(input)-1, target)
	t.Log(result)

	input = []int{1,2,3,5,6}
	target = 5
	result = BSLastSmaller(input, 0, len(input)-1, target)
	t.Log(result)
}

func TestGetResp(t *testing.T) {
	byteStr := "123 34 100 101 118 77 101 115 115 97 103 101 34 58 34 65 117 116 104 101 110 116 105 99 97 116 105 111 110 32 70 97 105 108 101 100 46 34 44 34 97 114 103 34 58 34 34 125"
	chars := strings.Fields(byteStr)
	byteArr := []byte{}
	for _, charInt := range chars {
		i, _ := strconv.Atoi(charInt)
		byteArr = append(byteArr, byte(i))
	}
	fmt.Println(string(byteArr))
}

func TestAtomic(t *testing.T) {
	atomicInt := int64(0)
	wg := &sync.WaitGroup{}
	go func(){
		for{
			fmt.Printf("%v\n", atomicInt)
		}
	}()
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
		    atomic.AddInt64(&atomicInt, 1)
		}()
	}
	wg.Wait()
}
