package main

import (
	"runtime"
	"sync"
)

func main() {
	var lock sync.Mutex

	lock.Lock()
	defer lock.Unlock()

	for range runtime.NumCPU() {
		go func() {
			var a int

			for {
				a++
			}
		}()
	}
}
