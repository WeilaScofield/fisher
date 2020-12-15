package main

import "errors"

type Stack interface {
	Size() int
	Nums() int
	Push(interface{}) error
	Pop() (interface{}, error)
	IsEmpty() bool
	IsFull() bool
	Elems() []interface{}
}

type ArrayStack struct {
	cap  int
	len  int
	data []interface{}
}

func NewArrayStack(c int) *ArrayStack {
	a := new(ArrayStack)
	a.cap = c
	a.len = 0
	a.data = make([]interface{}, a.len, a.cap)
	return a
}

func (s *ArrayStack) Size() int {
	return s.cap
}

func (s *ArrayStack) Nums() int {
	return s.len
}

func (s *ArrayStack) IsEmpty() bool {
	return s.len == 0
}

func (s *ArrayStack) IsFull() bool {
	return s.len == s.cap
}

func (s *ArrayStack) Push(value interface{}) error {
	if s.IsFull() {
		return errors.New("stack full")
	}

	s.data = append(s.data, value)
	s.len++

	return nil
}

func (s *ArrayStack) Pop() (value interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("stack empty")
	}

	value = s.data[s.len-1]
	s.data = s.data[:s.len-1]
	s.len = s.len - 1
	return value, nil
}

func (s *ArrayStack) Elems() []interface{} {
	return s.data
}
