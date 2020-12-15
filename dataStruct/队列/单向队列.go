package main

import (
	"errors"
	"fmt"
)

type queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (q *queue) add(val int) (err error) {
	if q.rear == q.maxSize {
		return errors.New("queue full")
	}

	q.rear++
	q.array[q.rear] = val
	return
}

func (q *queue) show() {
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]:%d\t", i, q.array[i])
	}
	fmt.Println()
}

func (q *queue) get() (val int, err error) {
	if q.front == q.rear {
		return -1, errors.New("queue empty")
	}

	q.front++

	return q.array[q.front], err
}

func main() {

	q := &queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	for i := 0; i < 5; i++ {
		err := q.add(i)
		if err != nil {
			fmt.Println(err)
		}
	}

	q.show()
	/*v, _ := q.get()
	fmt.Println(v)*/
	//q.show()
	/*v, _ = q.get()
	fmt.Println(v)*/
	//q.show()
	/*v, _ = q.get()
	fmt.Println(v)*/
}
