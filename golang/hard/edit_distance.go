package hard

import (
	"github.com/herlegs/programQ/golang/util"
)

func EditDistanceBruteForce(n, m string) int {
	if n == m {
		return 0
	}
	if n == "" || m == "" {
		return util.MaxInt(len(n), len(m))
	}
	//same
	if n[0] == m[0] {
		return EditDistanceBruteForce(n[1:], m[1:])
	}
	//insert in n
	insertDis := 1 + EditDistanceBruteForce(n, m[1:])
	//remove
	removeDis := 1 + EditDistanceBruteForce(n[1:], m)
	//replace
	replaceDis := 1 + EditDistanceBruteForce(n[1:], m[1:])

	return util.MinInt(util.MaxInt(insertDis, removeDis), replaceDis)
}

func EditDistanceRecursive(n, m string) int {
	ln, lm := len(n), len(m)
	result := make([][]int, ln+1)
	for i := 0; i <= ln; i++ {
		result[i] = make([]int, lm+1)
		for j := 0; j <= lm; j++ {
			result[i][j] = -1
		}
	}
	return editDistanceRecursive([]rune(n), []rune(m), ln, lm, result)
}

/*
result[i][j] store edit distance of n.substr(0,i) and m.substr(0,j)
*/
func editDistanceRecursive(nRunes, mRunes []rune, nIdx, mIdx int, result [][]int) int {
	if nIdx == 0 || mIdx == 0 {
		return nIdx + mIdx
	}
	if result[nIdx][mIdx] != -1 {
		return result[nIdx][mIdx]
	}
	//same last character
	if nRunes[nIdx-1] == mRunes[mIdx-1] {
		distance := editDistanceRecursive(nRunes, mRunes, nIdx-1, mIdx-1, result)
		result[nIdx][mIdx] = distance
		return distance
	}
	//insert character in n
	insertDistance := 1 + editDistanceRecursive(nRunes, mRunes, nIdx, mIdx-1, result)
	//remove character in n
	removeDistance := 1 + editDistanceRecursive(nRunes, mRunes, nIdx-1, mIdx, result)
	//replace character in n
	replaceDistance := 1 + editDistanceRecursive(nRunes, mRunes, nIdx-1, mIdx-1, result)

	minDistance := min3(insertDistance, removeDistance, replaceDistance)
	result[nIdx][mIdx] = minDistance

	return minDistance
}

func EditDistanceIterative(n, m string) int {
	ln, lm := len(n), len(m)
	nRunes, mRunes := []rune(n), []rune(m)
	result := make([][]int, ln+1)
	for i := 0; i <= ln; i++ {
		result[i] = make([]int, lm+1)
	}
	for i := 0; i <= ln; i++ {
		result[i][0] = i
	}
	for i := 0; i <= lm; i++ {
		result[0][i] = i
	}
	for i := 1; i <= ln; i++ {
		for j := 1; j <= lm; j++ {
			if nRunes[i-1] == mRunes[j-1] {
				result[i][j] = result[i-1][j-1]
				continue
			}
			result[i][j] = 1 + min3(result[i-1][j-1], result[i-1][j], result[i][j-1])
		}
	}
	return result[ln][lm]
}

func EditDistanceIterativeMemoryOptimize(n, m string) int {
	ln, lm := len(n), len(m)
	nRunes, mRunes := []rune(n), []rune(m)
	preRow, preCol := make([]int, lm+1), make([]int, ln+1)
	for i := 0; i <= lm; i++ {
		preRow[i] = i
	}
	for i := 0; i <= ln; i++ {
		preCol[i] = i
	}
	nextRow := make([]int, lm+1)
	for i := 1; i <= ln; i++ {
		nextRow[0] = preCol[i]
		for j := 1; j <= lm; j++ {
			if nRunes[i-1] == mRunes[j-1] {
				nextRow[j] = preRow[j-1]
				continue
			}
			nextRow[j] = 1 + min3(preRow[j-1], preRow[j], nextRow[j-1])
		}
		preRow, nextRow = nextRow, preRow
	}
	return preRow[lm]
}

func min3(a, b, c int) int {
	return util.MinInts(a, b, c)
}
