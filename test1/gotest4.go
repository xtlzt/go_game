package test1

import (
	"sync"
)

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		println("2222")
		instance = &singleton{}
	})
	return instance
}
