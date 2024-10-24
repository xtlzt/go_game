package test1

import (
	"fmt"
	"sync"
)

// sync.Mutex

var (
	x6 int64

	wg9 sync.WaitGroup // 等待组

	m4 sync.Mutex // 互斥锁
)

// add 对全局变量x执行5000次加1操作
func addh() {
	for i := 0; i < 5000; i++ {
		m4.Lock() // 修改x前加锁
		x6 = x6 + 1
		m4.Unlock() // 改完解锁
	}
	wg9.Done()
}

func Run9() {
	wg9.Add(2)

	go addh()
	go addh()

	wg9.Wait()
	fmt.Println(x6)
}
