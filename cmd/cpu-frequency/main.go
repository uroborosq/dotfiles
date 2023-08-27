package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, _ := os.Open("/proc/cpuinfo")
	scanner := bufio.NewScanner(file)
	var maxFreq float64
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 6 && s[:7] == "cpu MHz" {
			splittedStr := strings.Fields(s)
			
			if value, _ := strconv.ParseFloat(splittedStr[3], 64); value > maxFreq {
				maxFreq = value
			}
		}
	}
	fmt.Printf("%.1fMHz", maxFreq)
}
