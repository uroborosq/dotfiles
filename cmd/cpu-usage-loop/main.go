package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getStats(stats string) (int, int) {
	var (
		total, work int
		values      = strings.Split(stats, " ")[2:]
	)

	for i := range len(values) {
		parsedValue, _ := strconv.Atoi(values[i])
		total += parsedValue

		if i < 3 {
			work += parsedValue
		}
	}

	return total, work
}

func main() {
	for {
		fil, _ := os.Open("/proc/stat")
		buf := make([]byte, 1)
		fil.Read(buf)
		builder := strings.Builder{}
		for buf[0] != 10 {
			fil.Read(buf)
			builder.WriteByte(buf[0])
		}
		totalFirst, workFirst := getStats(builder.String())
		fil.Close()
		time.Sleep(time.Second)
		fil, _ = os.Open("/proc/stat")
		fil.Seek(0, 0)
		fil.Read(buf)
		builder.Reset()
		for buf[0] != 10 {
			fil.Read(buf)
			builder.WriteByte(buf[0])
		}
		totalSecond, workSecond := getStats(builder.String())

		fmt.Printf("%.1f%%\n", 100*float32(workSecond-workFirst)/float32(totalSecond-totalFirst))
		fil.Close()
	}
}
