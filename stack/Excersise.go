package stack

func reverseString(word string) *stack {
	var revStk stack
	for _, v := range word {
		revStk.Push(v)
	}
	return &revStk
}

func isBalanced(expr string) bool {
	var opFl stack
	var isBal bool
	for _, v := range expr {
		switch v {
		case R_OPEN_SQUARE, R_OPEN_PRNTH, R_OPEN_BRACE, R_OPEN_ANGLR:
			opFl.Push(v)
		case R_CLOSE_SQUARE, R_CLOSE_PRNTH, R_CLOSE_BRACE, R_CLOSE_ANGLR:
			if opFl.sizeOff() != 0 {
				popped := opFl.Pop().getElement()
				if v == R_CLOSE_ANGLR && popped == R_OPEN_ANGLR ||
					v == R_CLOSE_SQUARE && popped == R_OPEN_SQUARE ||
					v == R_CLOSE_PRNTH && popped == R_OPEN_PRNTH ||
					v == R_CLOSE_BRACE && popped == R_OPEN_BRACE {
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
