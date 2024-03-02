package main

import (
	"fmt"
	"linux/pkg/hwmon"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func gpuTemp(ch chan string) {
	output, _ := exec.Command("nvidia-smi", "--query-gpu=temperature.gpu", "--format=csv,noheader").Output()
	ch <- string(output[:len(output)-1])
}

func floatToTemp(value float64) string {
	s := strconv.FormatFloat(value, 'f', 0, 64)
	return strings.Join([]string{s[:2], ".", s[2:3]}, "")
}

func main() {
	gpu := make(chan string)
	go gpuTemp(gpu)

	parser := hwmon.NewSensorParser()

	sensors, err := parser.Parse()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("%s°C %s°C %s°C\n", floatToTemp(sensors["k10temp"]["Tctl"]), floatToTemp(sensors["nvme"]["Sensor 2"]), <-gpu)

}
