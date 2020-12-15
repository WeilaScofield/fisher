/*10个线程分别产生1个0-100间的随机整数，要求线程各自输出这个随机数，
从持有最小数的线程开始，最后是持有最大数的线程，依次输出*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type sendV struct {
	waitc chan struct{}
	value int
}

func newSendV(v int) sendV {
	return sendV{value: v, waitc: make(chan struct{})}
}

func quickSort(in []sendV) {
	if len(in) < 2 {
		return
	}

	l, r, p := 0, len(in)-1, in[0].value
	for i := 1; i <= r; {
		if in[i].value < p {
			in[i], in[l] = in[l], in[i]
			l++
			i++
		} else {
			in[i], in[r] = in[r], in[i]
			r--
		}
	}

	quickSort(in[:l])
	quickSort(in[l+1:])
}

func genRand(wg1, wg2 *sync.WaitGroup, sendc chan sendV) {
	v := rand.Intn(100)
	sendv := newSendV(v)
	sendc <- sendv
	wg1.Done()

	select {
	case <-sendv.waitc:
		fmt.Println(v)
	}
	wg2.Done()
}

func main() {
	wg1 := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	sendc := make(chan sendV, 10)

	for i := 0; i < 3; i++ {
		wg1.Add(1)
		wg2.Add(1)
		go genRand(wg1, wg2, sendc)
	}
	wg1.Wait()
	close(sendc)

	data := make([]sendV, 10, 10)
	index := 0
	for v := range sendc {
		data[index] = v
		index++
	}

	quickSort(data)

	for _, v := range data {
		v.waitc <- struct{}{}
		time.Sleep(1 * time.Second)
	}
	wg2.Wait()
}
