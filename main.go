package main

import (
	hashmap "GoDS/hash_map"
	"GoDS/linkedlist"
	"GoDS/queue"
	"GoDS/stack"
	"fmt"
)

func main() {
	pword := hashmap.PurifyString("               9sjldkbfkjsf pewytopwqytopqyt ,nbnm cdvdxv")
	fmt.Printf("%c\n", hashmap.GetFirstElement(pword, 1))
	stack.QueStk()
	stack.Exec()
	linkedlist.Exec()
	queue.Exec()
}
