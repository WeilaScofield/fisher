package main

import "fmt"

func main() {
	h := NewHead()
	n1 := NewNode(1)
	n2 := NewNode(5)
	n3 := NewNode(10)
	n4 := NewNode(3)
	h.Add(n1)
	h.Add(n2)
	h.Add(n3)
	h.Insert(n4)
	err := h.Del(4)
	if err != nil {
		fmt.Println(err)
	}
	h.Print()
	fmt.Println(h.Len())
}
