package main

import "fmt"

type A struct {
	b *B
}

type B struct {
	val int
}

func main() {
	a := A{
		b: &B{
			1,
		},
	}
	fmt.Println(*a.b)
	changeB(a)
	fmt.Println(*a.b)
}

func changeB(a A) {
	b := &B{
		2,
	}
	a.b = b
	fmt.Println(*a.b)
}
