package medium

func canWin(s string) bool {
	return canWinRecur(s, map[string]bool{})
}

func canWinRecur(s string, cache map[string]bool) bool {
	if win, exist := cache[s]; exist {
		return win
	}
	win := false
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] == '+' && s[i-1] == '+' {
			next := s[:i-1] + "--" + s[i+1:]
			if !canWinRecur(next, cache) {
				//fmt.Printf("next can't win:%v, so %v win\n", next, s)
				win = true
				break
			}
		}
	}
	cache[s] = win
	return win
}

func canLose(s string) bool {
	return canLoseRecur(s, map[string]bool{})
}

func canLoseRecur(s string, cache map[string]bool) bool {
	if win, exist := cache[s]; exist {
		return win
	}
	lose := false
	canFlip := false
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] == '+' && s[i-1] == '+' {
			canFlip = true
			next := s[:i-1] + "--" + s[i+1:]
			if !canLoseRecur(next, cache) {
				lose = true
				break
			}
		}
	}
	if !canFlip {
		lose = true
	}
	cache[s] = lose
	return lose
}
