package stack

func reverseString(word string) *Stack[rune] {
	var revStk Stack[rune]
	for _, v := range word {
		revStk.Push(v)
	}
	return &revStk
}

func isBalanced(expr string) bool {
	var opFl Stack[rune]
	var isBal bool
	for _, v := range expr {
		switch v {
		case '[', '(', '{', '<':
			opFl.Push(v)
		case ']', ')', '}', '>':
			if opFl.sizeOff() != 0 {
				popped := opFl.Pop().getElement()
				if v == '>' && popped == '<' ||
					v == ']' && popped == '[' ||
					v == ')' && popped == '(' ||
					v == '}' && popped == '{' {
					isBal = true
				} else {
					isBal = false
				}
			} else {
				isBal = false
			}
		default:
			continue
		}
	}
	return opFl.sizeOff() == 0 && isBal
}
