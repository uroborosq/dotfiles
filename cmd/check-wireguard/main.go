package main

import (
	"fmt"
	"github.com/prometheus/procfs"
	"time"
)

func main() {
	start := time.Now()

	fs, _ := procfs.NewFS("/proc")
	stats, _ := fs.Stat()
	fmt.Println(stats.CPUTotal.Nice)
	// _, err := net.InterfaceByName("peer2")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// file, _ := os.Open("/proc/net/dev")
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	if strings.Index(scanner.Text(), "peer2") != -1 {
	// 		fmt.Println("ok")
	// 	}
	// }

	fmt.Println(time.Since(start).Microseconds())
}
