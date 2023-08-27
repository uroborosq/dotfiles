package cpu

import (
	"bytes"
	"os"
	"strings"
	"time"
)

func toInt(s []byte) int {
	n := 0
	size := len(s)
	for i := 0; i < size; i++ {
		n = n*10 + int(s[i] - '0')
	}
	return n
}


func getStats(str []byte) (int, int) {
	valuesOfFirstTime := strings.Split(str, " ")[2:]
	total := 0
	work := 0
	valuesCount := len(valuesOfFirstTime)
	for i := 0; i < valuesCount; i++ {
		parsedValue := toInt(valuesOfFirstTime[i])
		total += parsedValue

		if i < 3 {
			work += parsedValue
		}
	}
	return total, work
}

func readLine(file *os.File) []byte {
	buffer := make([]byte, 20)
	builder := bytes.Buffer{}
	for {
		file.Read(buffer)
		for i := 0; i < 20; i++ {
			if buffer[i] == 10 {
				builder.Write(buffer[:i])
				break
			}
		}
	}
	return builder.Bytes()
}


func GetValue() string {
	file, _ := os.Open("/proc/stat")
	totalFirst, workFirst := getStats(readLine(file))
	file.Close()
	time.Sleep(time.Second)
	file, _ = os.Open("/proc/stat")
	totalSecond, workSecond := getStats(readLine(file))
	file.Close()
	return :
}