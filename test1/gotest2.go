package test1

import (
	"fmt"
	"strconv"
	"sync"
)

// 并发安全的map
var m1 = sync.Map{}

func Run2() {
	GetInstance()
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m1.Store(key, n)         // 存储key-value
			value, _ := m1.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
