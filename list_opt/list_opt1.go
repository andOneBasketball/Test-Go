package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	l.PushBack("a")
	printList(l) // a

	l.PushBack("b")
	printList(l) // a b

	l.PushFront("c")
	printList(l) // c a b

	fmt.Println(l.Front().Value) // c
	fmt.Println(l.Back().Value)  // b
	fmt.Println(l.Len())         // 3

	l.MoveToBack(l.Front())
	printList(l) // a b c

	l.MoveToFront(l.Back())
	printList(l) // c a b

	l.Remove(l.Back())
	printList(l) // c a
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
