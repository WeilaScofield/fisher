package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// when：多个goroutines等待、1个goroutine通知事件发生。

const (
	placeholder = "place"
	target      = 7
	wide        = 10
)

var wg = &sync.WaitGroup{}

type limitMap struct {
	lock       sync.Mutex
	cond       sync.Cond
	flushCount int
	m          map[string]int
}

func main() {
	rand.Seed(time.Now().Unix())

	wg.Add(1)

	ls := newLimitSlice()
	go ls.autoFetch()
	go ls.autoFlush()
	go ls.autoAdd()

	wg.Wait()
}

func newLimitSlice() *limitMap {
	ls := &limitMap{
		m: make(map[string]int),
	}
	ls.cond.L = &ls.lock
	return ls
}

func (ls *limitMap) autoAdd() {
	time.Sleep(2 * time.Second)

	ls.lock.Lock()
	k := rand.Intn(wide)
	ls.m[placeholder] = k
	ls.lock.Unlock()

	ls.cond.Broadcast()
}

func (ls *limitMap) autoFlush() {
	for {
		time.Sleep(3 * time.Second)
		ls.lock.Lock()

		fmt.Println("to flush")
		ls.m = make(map[string]int)
		ls.flushCount++
		ls.lock.Unlock()
		ls.autoAdd()
	}
}

func (ls *limitMap) autoFetch() {
	for {
		ls.lock.Lock()
		for len(ls.m) == 0 {
			fmt.Println("fetch wait for cond")
			ls.cond.Wait()
		}
		v := ls.m[placeholder]

		fmt.Println("to print:", v)

		if v == target {
			fmt.Printf("flush count: %v\n", ls.flushCount)
			wg.Done()
			ls.lock.Unlock()
			return
		}

		ls.lock.Unlock()

		time.Sleep(1 * time.Second)
	}
}
