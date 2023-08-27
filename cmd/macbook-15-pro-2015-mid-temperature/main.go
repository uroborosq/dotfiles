package main

import (
	"fmt"
	"os"
	"strconv"
)


func averageCpuTemp(paths []string) string {
	sum := 0
	counter := 0
	buffer := make([]byte, 5)
	for _, path := range paths {
		file, _ := os.Open(path)
		
		file.Read(buffer)
		temp, _ := strconv.Atoi(string(buffer))
		sum += temp
		counter++
	}
	average := float64(sum) / float64(counter) / 1000
	return fmt.Sprintf("%0.1f ℃", average)
}

func main() {
	cpuPaths := []string{"/sys/class/hwmon/hwmon5/temp1_input", "/sys/class/hwmon/hwmon5/temp2_input", "/sys/class/hwmon/hwmon5/temp3_input", "/sys/class/hwmon/hwmon5/temp4_input"}
	fmt.Println(averageCpuTemp(cpuPaths))
}