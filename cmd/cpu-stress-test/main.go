package main

import (
	"runtime"
	"sync"
)

func main() {
	coreNumber := runtime.NumCPU()
	var lol sync.Mutex
	lol.Lock()
	for i := 0; i < coreNumber; i++ {
		go func() {
			a := 0
			for {
				a++
			}
		}()
	}
	lol.Lock()
}
