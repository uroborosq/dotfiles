package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"linux/pkg/hwmon"
)

func main() {
	parser := hwmon.NewSensorParser()

	sensors, err := parser.Parse()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("%s°C %s°C", floatToTemp(sensors["coretemp"]["Package id 0"]), floatToTemp(sensors["nvme"]["Composite"]))
}

func floatToTemp(value float64) string {
	s := strconv.FormatFloat(value, 'f', 0, 64)

	return strings.Join([]string{s[:2], ".", s[2:3]}, "")
}
