/*场景：在一个高并发的web服务器中，要限制IP的频繁访问。
现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]struct{}
	mu       *sync.Mutex
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]struct{}), mu: &sync.Mutex{}}
}

//判断IP是否存在
func (o *Ban) visit(ip string) bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = struct{}{}
	go o.invalidAfter3Min(ip)
	return false
}

// 3分钟后ip失效, 从map中移除. 因此ip再次访问时便可正常访问
func (o *Ban) invalidAfter3Min(ip string) {
	time.Sleep(time.Minute * 3)
	o.mu.Lock()
	visitIPs := o.visitIPs
	delete(visitIPs, ip)
	o.visitIPs = visitIPs
	o.mu.Unlock()
}

func main1() {
	var success int64
	ban := NewBan()
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			ipEnd := j // 闭包
			go func() {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", ipEnd)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}()
		}
	}
	wg.Wait()
	fmt.Println("success:", success)
}
