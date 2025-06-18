package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"unsafe"

	"linux/pkg/hp"
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
			freqBytes := hp.P(os.ReadFile(fmt.Sprintf(cpuFrequencySysPath, i)))
			freqStr := unsafe.String(unsafe.SliceData(freqBytes), len(freqBytes))
			freq := hp.P(strconv.ParseInt(freqStr[:len(freqBytes)-1], 10, 64))

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
