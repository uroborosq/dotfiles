package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"linux/pkg/logger"
)

const statFile = "/proc/stat"

func getStats(str string) (int, int) {
	valuesOfFirstTime := strings.Split(str, " ")[2:]
	total := 0
	work := 0
	valuesCount := len(valuesOfFirstTime)
	for i := range valuesCount {
		parsedValue, _ := strconv.Atoi(valuesOfFirstTime[i])
		total += parsedValue

		if i < 3 {
			work += parsedValue
		}
	}

	return total, work
}

type cpuStats struct {
	work  int
	total int
}

func readFile(path string) (cpuStats, error) {
	var stats cpuStats

	reader, err := os.Open(path)
	if err != nil {
		return stats, fmt.Errorf("can't open %s: %w", path, err)
	}

	buf := make([]byte, 1)
	builder := strings.Builder{}

	for buf[0] != 10 {
		if count, readErr := reader.Read(buf); readErr != nil {
			return stats, fmt.Errorf("can't read bytes from %s: %w", statFile, readErr)
		} else if count != 1 {
			return stats, fmt.Errorf("expected to read 1 byte, read %d", count)
		}

		builder.WriteByte(buf[0])
	}

	stats.total, stats.work = getStats(builder.String())

	return stats, nil
}

func main() {
	execType := flag.String("type", "once", "once or loop")

	flag.Parse()

	for {
		first, err := readFile(statFile)
		if err != nil {
			logger.Fatalf("%s", err.Error())
		}

		time.Sleep(time.Second)

		second, err := readFile(statFile)
		if err != nil {
			logger.Fatalf("%s", err.Error())
		}

		fmt.Printf("%.1f%%\n", 100*float32(second.work-first.work)/float32(second.total-first.total))

		if *execType == "once" {
			break
		}
	}
}
