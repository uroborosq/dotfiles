package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getStats(str string) (int, int) {
	valuesOfFirstTime := strings.Split(str, " ")[2:]
	total := 0
	work := 0
	valuesCount := len(valuesOfFirstTime)
	for i := 0; i < valuesCount; i++ {
		parsedValue, _ := strconv.Atoi(valuesOfFirstTime[i])
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
