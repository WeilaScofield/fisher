package main

import "fmt"

type boy struct {
	no   int
	next *boy
}

func addBoy(n int) *boy {
	first := &boy{}
	temp := first

	if n < 1 {
		fmt.Println("too less kid")
	}

	for i := 1; i <= n; i++ {
		b := &boy{
			no: i,
		}
		if i == 1 {
			first = b
			temp = b
			temp.next = first
		} else {
			temp.next = b
			temp.next.next = first
		}
		temp = temp.next
	}
	return first
}

func displayBoy(first *boy) {
	temp := first
	if first.next == nil {
		return
	}
	for {
		fmt.Printf("boy[%d]-->", temp.no)
		if temp.next == first {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

func josephuGame(first *boy, startNo, countNo int) {
	tail := first
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}

	if startNo > tail.no {
		fmt.Println("invalid startNo")
		return
	}

	for i := 0; i < startNo-1; i++ {
		first = first.next
		tail = tail.next
	}

	for {
		for i := 0; i < countNo-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("boy[%d] out of circle\n", first.no)
		first = first.next
		tail.next = first
		if first == tail {
			fmt.Printf("boy[%d] out of circle\n", first.no)
			break
		}
	}

}

func main() {
	first := addBoy(7)
	displayBoy(first)
	josephuGame(first, 2, 3)
}
