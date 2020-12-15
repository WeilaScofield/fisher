package main

import "fmt"

func fibonacci(n int) {
	a := 0
	b := 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}
	fmt.Println(b)
}

func main() {
	fibonacci(5)
}
