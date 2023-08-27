package main

import (
	"fmt"
	"linux/pkg/temperature"
)

func main() {
	cpuPath := "/sys/devices/platform/thinkpad_hwmon/hwmon/hwmon3/temp1_input"
	gpuPath := "/sys/devices/platform/thinkpad_hwmon/hwmon/hwmon3/temp2_input"

	cpuFirst, cpuSecond := make(chan []byte), make(chan byte)
	gpuFirst, gpuSecond := make(chan []byte), make(chan byte)

	go temperature.Get(cpuPath, cpuFirst, cpuSecond)
	go temperature.Get(gpuPath, gpuFirst, gpuSecond)

	fmt.Printf("%s.%c°C ", <-cpuFirst, <-cpuSecond)
	fmt.Printf("%s.%c°C \n", <-gpuFirst, <-gpuSecond)
}
