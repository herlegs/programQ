package hard

import "sort"

//given a list of chars (as a string), and target word length, output all possible word combination
func PickWords(charList string, wordLen int) []string {
	subsets := Subset(charList, wordLen)
	words := []string{}
	for _, subset := range subsets {
		words = append(words, Permutation(subset)...)
	}
	return words
}

//given a list of chars (as a string), output all unique subset of chars (as string)
func Subset(list string, subsetSize int) []string {
	chars := []rune(list)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	if subsetSize >= len(list) {
		return []string{list}
	}
	result := []string{}
	subsetRecur(chars, 0, subsetSize, []rune{}, &result)
	return result
}

func subsetRecur(list []rune, currIndex, targetLen int, tmp []rune, result *[]string) {
	if len(tmp) == targetLen {
		*result = append(*result, string(tmp))
		return
	}
	for i:= currIndex; i < len(list); i++ {
		if i != currIndex && list[i] == list[i-1] {
			continue
		}
		tmp = append(tmp, list[i])
		subsetRecur(list, i+1, targetLen, tmp, result)
		tmp = tmp[:len(tmp)-1]
	}
}

//given a list of chars (as a string), output all unique permutation of the chars
func Permutation(str string) []string {
	chars := []rune(str)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
 	})
	result := []string{}
	permutationRecur(chars, 0, len(chars) - 1, &result)
	return result
}

func permutationRecur(list []rune, start, end int, result *[]string) {
	if start == end {
		*result = append(*result, string(list))
		return
	}
	for i:= start; i <= end; i++ {
		if i != start && list[i] == list[i-1] {
			continue
		}
		list[start], list[i] = list[i], list[start]
		permutationRecur(list, start+1, end, result)
		list[start], list[i] = list[i], list[start]
	}
}
