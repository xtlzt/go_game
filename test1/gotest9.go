package test1

import (
	"fmt"
	"time"
)

// demo2 通道误用导致的bug
func Demo2() {
	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		//time.Sleep(1 * time.Second)
		fmt.Println("99999999999999999")
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		return
	}
}
