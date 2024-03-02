package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"unsafe"
)

func main() {
	var maxFreq atomic.Int64
	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())
	for i := range runtime.NumCPU() {
		go func() {
			freqBytes, err := os.ReadFile(fmt.Sprintf("/sys/devices/system/cpu/cpu%d/cpufreq/scaling_cur_freq", i))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			freqStr := unsafe.String(unsafe.SliceData(freqBytes), len(freqBytes))
			freq, err := strconv.ParseInt(freqStr[:len(freqBytes)-1], 10, 64)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			for true {
				currentMaxFreq := maxFreq.Load()
				if freq > currentMaxFreq && maxFreq.CompareAndSwap(currentMaxFreq, freq) {
					break
				} else if freq <= currentMaxFreq {
					break
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%vMHz\n", maxFreq.Load()/1000)
}
