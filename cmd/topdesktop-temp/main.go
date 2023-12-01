package main

import (
	"fmt"
	"os"
	"os/exec"
)

func gpuTemp(ch chan string) {
	output, _ := exec.Command("nvidia-smi", "--query-gpu=temperature.gpu", "--format=csv,noheader").Output()
	ch <- string(output[:len(output)-1])
}

func diskTemp(ch chan string) {
	output, _ := os.ReadFile("/sys/class/hwmon/hwmon0/temp3_input")
	ch <- string(output[:2])+"."+string(output[2])
}

func cpuTemp(ch chan string) {
	output, _ := os.ReadFile("/sys/class/hwmon/hwmon1/temp1_input")
	ch <- string(output[:2])+"."+string(output[2])
}

func main() {
	gpu := make(chan string)
	cpu := make(chan string)
	disk := make(chan string)

	go gpuTemp(gpu)
	go diskTemp(disk)
	go cpuTemp(cpu)

	fmt.Printf("%s°C %s°C %s°C\n", <-cpu, <-disk, <-gpu)
}