package test1

import (
	"fmt"
	"sync"
)

var (
	x1 int64

	wg3 sync.WaitGroup // 等待组
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		x1 = x1 + 1
	}
	wg3.Done()
}

func Run6() {
	wg3.Add(2)

	go add()
	go add()

	wg3.Wait()
	fmt.Println(x1)
}
