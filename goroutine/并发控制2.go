//评测题目: 问题1：三个线程交替打印antantant...，一个打印a，一个打印n，一个打印t
package main

import (
	"fmt"
	"time"
)

type ant struct {
	v     string
	waitc chan struct{}
}

func main() {
	a1 := ant{"a", make(chan struct{})}
	a2 := ant{"n", make(chan struct{})}
	a3 := ant{"t", make(chan struct{})}

	go antPrint(a1, a2)
	go antPrint(a2, a3)
	go antPrint(a3, a1)

	a1.waitc <- struct{}{}

	select {}
}

func antPrint(a, b ant) {
	for {
		select {
		case <-a.waitc:
			fmt.Println(a.v)
			b.waitc <- struct{}{}
		}
		time.Sleep(time.Second)
	}
}
