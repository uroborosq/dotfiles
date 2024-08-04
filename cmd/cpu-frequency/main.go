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

const (
	multpler            = 1000
	cpuFrequencySysPath = "/sys/devices/system/cpu/cpu%d/cpufreq/scaling_cur_freq"
)

func main() {
	var (
		maxFreq   atomic.Int64
		waitGroup sync.WaitGroup
	)

	waitGroup.Add(runtime.NumCPU())

	for i := range runtime.NumCPU() {
		go func() {
			freqBytes, err := os.ReadFile(fmt.Sprintf(cpuFrequencySysPath, i))
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

			for {
				currentMaxFreq := maxFreq.Load()
				if freq > currentMaxFreq && maxFreq.CompareAndSwap(currentMaxFreq, freq) {
					break
				} else if freq <= currentMaxFreq {
					break
				}
			}
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	kHZ := maxFreq.Load()
	fmt.Printf("%vMHz\n", kHZ/multpler)
}
