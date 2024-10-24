package test1

import (
	"fmt"
	"sync"
	"time"
)

var (
	x       int64
	wg1     sync.WaitGroup
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	mutex.Unlock()                    // 解互斥锁
	wg1.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()                 // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutex.Unlock()               // 释放互斥锁
	wg1.Done()
}

// writeWithLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.Unlock()                  // 释放写锁
	wg1.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwMutex.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutex.RUnlock()            // 释放读锁
	wg1.Done()
}

func Do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg1.Add(1)
		go wf()
	}

	//  rc个并发读操作
	for i := 0; i < rc; i++ {
		wg1.Add(1)
		go rf()
	}

	wg1.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)

}

func DoY() {
	Do(writeWithLock, readWithLock, 10, 1000) // x:10 cost:1.466500951s

	// 使用读写互斥锁，10并发写，1000并发读
	Do(writeWithRWLock, readWithRWLock, 10, 1000) // x:10 cost:117.207592ms
}
