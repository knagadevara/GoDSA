package stack

type stkElm struct {
	element        rune
	abvElm, bloElm *stkElm
}

type stack struct {
	size      int
	firstElem *stkElm
}

/*
Methods
push -> puts an item on the stack
pop -> removes the first item on Stack
peek -> returns the first item on stack without removing it
isEmpty -> checks if stack contains anything and returns bool
sizeOff -> returns number of elements in stack.
reverse -> Reversing a String
*/
