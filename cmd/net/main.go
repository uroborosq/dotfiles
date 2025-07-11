package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readStats(scanner *bufio.Scanner) (int64, int64) {
	send := int64(0)
	received := int64(0)

	for i := 0; scanner.Scan(); i++ {
		if i < 2 {
			continue
		}

		splittedString := strings.Fields(scanner.Text())
		if strings.HasPrefix(splittedString[0], "enp") || strings.HasPrefix(splittedString[0], "wlp") || strings.HasPrefix(splittedString[0], "wlo") {
			tmp, _ := strconv.ParseInt(splittedString[1], 10, 64)
			received += tmp
			tmp, _ = strconv.ParseInt(splittedString[9], 10, 64)
			send += tmp
		}
	}

	return received, send
}

func main() {
	file, _ := os.Open("/proc/net/dev")
	scanner := bufio.NewScanner(file)

	received, send := readStats(scanner)

	file.Close()
	time.Sleep(1 * time.Second)

	file, _ = os.Open("/proc/net/dev")
	scanner = bufio.NewScanner(file)

	receivedSecond, sendSecond := readStats(scanner)
	received -= receivedSecond
	send -= sendSecond

	file.Close()

	sendHumanView := 8 * float64(-send)
	receivedHumanView := 8 * float64(-received)

	measurePoints := []string{
		"b/s", "Kib/s", "Mib/s", "Gib/s",
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
