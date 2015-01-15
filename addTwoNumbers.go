package main

// You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
import (
	"container/list"
	"fmt"
)

func main() {
	l1 := list.New()
	l1.PushBack(4)
	l1.PushBack(5)
	l1.PushBack(2)

	l2 := list.New()
	l2.PushBack(7)
	l2.PushBack(3)

	l3 := list.New()
	l3 = addTwoNumbers(l1, l2)

	for e := l3.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
func addTwoNumbers(l1 *list.List, l2 *list.List) (l3 *list.List) {
	carry := 0
	l4 := list.New()
	e1 := l1.Front()
	e2 := l2.Front()
	for {
		sum := carry
		if e1 != nil {
			sum += e1.Value.(int)
			e1 = e1.Next()
		}
		if e2 != nil {
			sum += e2.Value.(int)
			e2 = e2.Next()
		}
		l4.PushBack(sum % 10)
		carry = sum / 10

		if e1 == nil && e2 == nil && carry == 0 {
			break
		}
	}
	return l4
}
