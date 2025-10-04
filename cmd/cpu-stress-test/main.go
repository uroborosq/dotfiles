package main

import (
	"runtime"
	"sync/atomic"
)

func main() {
	for range runtime.NumCPU() {
		go func() {
			var a atomic.Int64

			for {
				a.Add(1)
			}
		}()
	}

	select {}
}
