package test1

import (
	"fmt"
	"sync"
)

// demo1 通道误用导致的bug
func Demo5() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			defer wg.Done()
			for {
				//task, ok := <-ch
				//// 这里假设对接收的数据执行某些操作
				//fmt.Println(task)
				//if !ok {
				//	break
				//}
				select {
				case task, ok := <-ch: // 从通道接收
					if !ok { // 如果通道已关闭
						return // 退出 goroutine
					}
					// 处理接收到的数据
					fmt.Println(task)
				}
			}

		}()
	}
	wg.Wait()
}
