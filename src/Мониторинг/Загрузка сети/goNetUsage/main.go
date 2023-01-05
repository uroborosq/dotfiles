package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open("/proc/net/dev")
	scanner := bufio.NewScanner(file)

	send := int64(0)
	received := int64(0)

	for i := 0; scanner.Scan(); i++ {
		if i < 2 {
			continue
		}

		splittedString := strings.Fields(scanner.Text())

		if splittedString[0] == "wlo1:" || splittedString[0] == "enp3s0f3u1u1:" {
			tmp, _ := strconv.ParseInt(splittedString[1], 10, 64)
			received += tmp
			tmp, _ = strconv.ParseInt(splittedString[9], 10, 64)
			send += tmp
		}
	}
	file.Close()
	time.Sleep(1 * time.Second)

	file, _ = os.Open("/proc/net/dev")
	scanner = bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		if i < 2 {
			continue
		}

		splittedString := strings.Fields(scanner.Text())

		if splittedString[0] == "wlo1:" || splittedString[0] == "enp3s0f3u1u1:" {
			tmp, _ := strconv.ParseInt(splittedString[1], 10, 64)
			received -= tmp
			tmp, _ = strconv.ParseInt(splittedString[9], 10, 64)
			send -= tmp
		}
	}
	file.Close()

	sendHumanView := float64(-send)
	receivedHumanView := float64(-received)

	measurePoints := []string{
		"B/s", "KiB/s", "MiB/s", "GiB/s",
	}

	sendMeasure := 0
	receivedMeasure := 0

	for sendHumanView > 1024 {
		sendHumanView /= 1024
		sendMeasure++
	}

	for receivedHumanView > 1024 {
		receivedHumanView /= 1024
		receivedMeasure++
	}

	fmt.Printf("%.1f%s %.1f%s", receivedHumanView, measurePoints[receivedMeasure], sendHumanView, measurePoints[sendMeasure])
}
