package main

import (
	"fmt"
	"os"
	"runtime"
)

func get(path string, first chan []byte, second chan byte) {
	data := make([]byte, 6)
	sensor, _ := os.Open(path)
	sensor.Read(data)
	first <- data[:2]
	second <- data[2]
}

func main() {
	runtime.GOMAXPROCS(8)
	firstFirst := make(chan []byte)
	firstSecond := make(chan byte)
	secondFirst := make(chan []byte)
	secondSecond := make(chan byte)
	thirdFirst := make(chan []byte)
	thirdSecond := make(chan byte)
	go get("/sys/module/k10temp/drivers/pci:k10temp/0000:00:18.3/hwmon/hwmon5/temp1_input", firstFirst, firstSecond)
	go get("/sys/module/amdgpu/drivers/pci:amdgpu/module/drivers/pci:amdgpu/0000:03:00.0/hwmon/hwmon4/temp1_input", secondFirst, secondSecond)
	go get("/sys/module/nvme/drivers/pci:nvme/0000:02:00.0/nvme/nvme0/hwmon3/temp3_input", thirdFirst, thirdSecond)

	fmt.Printf("%s.%c°C ", <-firstFirst, <-firstSecond)
	fmt.Printf("%s.%c°C ", <-secondFirst, <-secondSecond)
	fmt.Printf("%s.%c°C\n", <-thirdFirst, <-thirdSecond)
}
