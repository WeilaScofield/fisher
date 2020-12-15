package main

import "fmt"

func main() {
	var s Stack = NewArrayStack(10)
	err := s.Push(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.Elems())
}
