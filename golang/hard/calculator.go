package hard

import "container/list"

func calculate(s string) int {
	processed := list.New()
	opStack := list.New()
	digitBuffer := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if isDigit(ch) {
			digitBuffer = append(digitBuffer, ch)
			continue
		}
		if !isOperand(ch) {
			continue
		}
		if len(digitBuffer) != 0 {
			processed.PushBack(toDigit(digitBuffer))
		}
		digitBuffer = []byte{}
		if ch == ')' {
			//pop until '('
			for opStack.Back().Value.(byte) != '(' {
				e := opStack.Back()
				opStack.Remove(e)
				//processed.PushBack(e.Value.(byte))
				doOp(e.Value.(byte), processed)
			}
			opStack.Remove(opStack.Back())

		} else if opStack.Len() == 0 || !compareOp(opStack.Back().Value.(byte), ch) {
			opStack.PushBack(ch)
		} else {
			//pop out all higher rank operand
			for opStack.Len() != 0 && compareOp(opStack.Back().Value.(byte), ch) {
				e := opStack.Back()
				opStack.Remove(e)
				//processed.PushBack(e.Value.(byte))
				doOp(e.Value.(byte), processed)
			}
			opStack.PushBack(ch)
		}
	}
	if len(digitBuffer) != 0 {
		processed.PushBack(toDigit(digitBuffer))
	}
	for opStack.Len() != 0 {
		e := opStack.Back()
		opStack.Remove(e)
		doOp(e.Value.(byte), processed)
	}

	return processed.Back().Value.(int)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func toDigit(bs []byte) int {
	num := 0
	for _, b := range bs {
		num = num*10 + int(b-'0')
	}
	return num
}

func isOperand(b byte) bool {
	if b == '+' || b == '-' || b == '*' || b == '/' || b == '(' || b == ')' {
		return true
	}
	return false
}

var OpMap = map[byte]int{
	'+': 1, '-': 1,
	'*': 2, '/': 2,
}

//return true (higher, should pop out) false (lower, should keep)
func compareOp(left, right byte) bool {
	if left == '(' || right == '(' {
		return false
	}
	return OpMap[left] >= OpMap[right]
}

func doOp(op byte, stack *list.List) {
	right := stack.Back()
	stack.Remove(right)
	left := stack.Back()
	stack.Remove(left)

	leftVal, rightVal := left.Value.(int), right.Value.(int)

	var result int
	switch op {
	case '+':
		result = leftVal + rightVal
	case '-':
		result = leftVal - rightVal
	case '*':
		result = leftVal * rightVal
	case '/':
		result = leftVal / rightVal
	}

	stack.PushBack(result)
}
