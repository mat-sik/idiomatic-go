package main

import "fmt"

func main() {
	list := &linkedList[int]{}
	list.add(10)
	list.add(20)
	list.add(30)

	list.insert(3, 0)
	list.insert(1, 0)
	list.insert(2, 1)
	list.insert(5, 3)
	list.insert(4, 3)
	list.insert(15, 6)

	list.insert(40, 9)

	list.insert(100, 100)

	fmt.Println("Index of 1:", list.index(1))
	fmt.Println("Index of 2:", list.index(2))
	fmt.Println("Index of 3:", list.index(3))
	fmt.Println("Index of 4:", list.index(4))
	fmt.Println("Index of 5:", list.index(5))
	fmt.Println("Index of 10:", list.index(10))
	fmt.Println("Index of 15:", list.index(15))
	fmt.Println("Index of 20:", list.index(20))
	fmt.Println("Index of 30:", list.index(30))
	fmt.Println("Index of 40:", list.index(40))
	fmt.Println("Index of 100:", list.index(100))
}

type listNode[T comparable] struct {
	value T
	next  *listNode[T]
}

type linkedList[T comparable] struct {
	head *listNode[T]
	tail *listNode[T]
}

func (l *linkedList[T]) add(value T) {
	node := &listNode[T]{value: value}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
}

func (l *linkedList[T]) insert(value T, idx int) {
	if l.head == nil {
		l.add(value)
		return
	}
	if idx == 0 {
		node := &listNode[T]{value: value}
		node.next = l.head
		l.head = node
		return
	}

	curr := l.head
	parentIdx := idx - 1
	for currIdx := 0; curr != l.tail && currIdx < parentIdx; curr, currIdx = curr.next, currIdx+1 {
	}

	node := &listNode[T]{value: value}
	node.next = curr.next
	curr.next = node
	if node.next == nil {
		l.tail = node
	}
}

func (l *linkedList[T]) index(value T) int {
	idx := 0
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.value == value {
			return idx
		}
		idx++
	}
	return -1
}
