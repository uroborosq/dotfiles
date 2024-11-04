package main

import (
	"context"
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

	<-context.Background().Done()
}
