package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("/proc/meminfo")
	bufReader := bufio.NewReader(file)
	memTotalStr, _ := bufReader.ReadString('\n')
	bufReader.ReadLine()
	memAvailableStr, _ := bufReader.ReadString('\n')
	for i := 0; i < 11; i++ {
		bufReader.ReadLine()
	}

	swapTotalStr, _ := bufReader.ReadString('\n')
	swapAvailableStr, _ := bufReader.ReadString('\n')

	memTotal, _ := strconv.ParseFloat(strings.Fields(memTotalStr)[1], 32)
	memAvailable, _ := strconv.ParseFloat(strings.Fields(memAvailableStr)[1], 32)
	swapTotal, _ := strconv.ParseFloat(strings.Fields(swapTotalStr)[1], 32)
	swapAvailable, _ := strconv.ParseFloat(strings.Fields(swapAvailableStr)[1], 32)

	fmt.Printf("%.1fGiB %.1fGiB\n", (memTotal-memAvailable) / 1024 / 1024, (swapTotal-swapAvailable) / 1024 / 1024)
}
