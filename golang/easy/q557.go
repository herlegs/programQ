package easy

import "fmt"

/*
https://leetcode.com/problems/reverse-words-in-a-string-iii

Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
Note: In the string, each word is separated by single space and there will not be any extra space in the string.
 */
func ReverseWords(s string) string {
	fmt.Println(s)
	array := []byte(s)
	n := len(array)
	for i := 0; i < n; i++ {
		j := i
		for j < n && array[j] != ' '{
			j++
		}
		reverse(array, i, j - 1)
		i = j
	}
	return string(array)
}

func reverse(array []byte, s, e int) {
	for i,j := s,e; i < j; i,j = i+1,j-1 {
		tmp := array[i]
		array[i] = array[j]
		array[j] = tmp
	}
}


