package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"

	"linux/pkg/hp"
)

func main() {
	err := pprof.StartCPUProfile(hp.P(os.OpenFile("cpu.prof", os.O_CREATE|os.O_WRONLY, 0644)))
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	content, _ := os.ReadFile("/proc/meminfo")
	lines := strings.Split(string(content), "\n")

	memTotal, _ := strconv.ParseFloat(strings.Fields(lines[0])[1], 32)
	memAvailable, _ := strconv.ParseFloat(strings.Fields(lines[2])[1], 32)
	swapTotal, _ := strconv.ParseFloat(strings.Fields(lines[14])[1], 32)
	swapAvailable, _ := strconv.ParseFloat(strings.Fields(lines[15])[1], 32)

	fmt.Printf("%.1fGiB %.1fGiB\n", (memTotal-memAvailable)/1024/1024, (swapTotal-swapAvailable)/1024/1024)
}
