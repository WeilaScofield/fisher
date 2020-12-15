package main

import (
	"errors"
	"fmt"
)

type ringQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

func (r *ringQueue) push(val int) (err error) {
	if r.isFull() {
		return errors.New("ringQueue full")
	}
	r.array[r.tail] = val
	r.tail = (r.tail + 1) % r.maxSize
	return
}

func (r *ringQueue) pop() (val int, err error) {
	if r.isEmpty() {
		return 0, errors.New("ringQueue empty")
	}
	val = r.array[r.tail]
	r.head = (r.head + 1) % r.maxSize
	return
}

func (r *ringQueue) isFull() bool {
	return (r.tail+1)%r.maxSize == r.head
}

func (r *ringQueue) isEmpty() bool {
	return r.tail == r.head
}

func (r *ringQueue) size() int {
	return (r.tail + r.maxSize - r.head) % r.maxSize
}

func (r *ringQueue) display() {
	if r.isEmpty() {
		fmt.Println("ringQueue empty")
	}

	tempHead := r.head
	for i := 0; i < r.size(); i++ {
		fmt.Printf("queue[%d]:%d\t", tempHead, r.array[tempHead])
		tempHead = (tempHead + 1) % r.maxSize
	}
	fmt.Println()
}

func main() {
	r := &ringQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	r.push(1)
	r.push(2)
	r.push(3)
	r.push(4)

	r.display()

	r.pop()
	r.pop()
	r.display()

	r.push(5)
	r.push(6)
	r.display()
}
