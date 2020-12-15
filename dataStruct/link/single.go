package main

import (
	"errors"
	"fmt"
)

type SingleLink struct {
	value int
	next  *SingleLink
}

func NewHead() *SingleLink {
	return &SingleLink{}
}

func NewNode(v int) *SingleLink {
	return &SingleLink{value: v}
}

func (l *SingleLink) IsEmpty() bool {
	return l.next == nil
}

func (l *SingleLink) Print() {
	temp := l
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
		fmt.Printf("%v->", temp.value)
	}
	fmt.Println("nil")
}

func (l *SingleLink) Add(newNode *SingleLink) {
	temp := l
	for {
		if temp.next == nil {
			temp.next = newNode
			break
		}
		temp = temp.next
	}
}

func (l *SingleLink) Del(value int) error {
	if l.IsEmpty() {
		return errors.New("link empty")
	}
	temp := l
	for {
		if temp.next == nil {
			return errors.New("find no such node")
		}
		if temp.next.value == value {
			temp.next = temp.next.next
		}
		temp = temp.next
	}
}

func (l *SingleLink) Insert(newNode *SingleLink) {
	if l.IsEmpty() {
		l.next = newNode
		return
	}
	temp := l
	for {
		if temp.next == nil {
			temp.next = newNode
			break
		}
		if newNode.value <= temp.next.value {
			newNode.next = temp.next
			temp.next = newNode
			break
		}
		temp = temp.next
	}
}

func (l *SingleLink) Len() int {
	nums := 0
	temp := l
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
		nums++
	}
	return nums
}
