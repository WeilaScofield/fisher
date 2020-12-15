package main

import "fmt"

func fibonacci2(n int) int {
	if n < 3 {
		return 1
	}
	n = fibonacci2(n-1) + fibonacci2(n-2)
	return n
}

func main() {
	fmt.Println(fibonacci2(6))
}
