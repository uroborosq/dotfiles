package main

import (
	"fmt"
	"time"

	"github.com/prometheus-community/pro-bing"
)

func main() {
	
	for {
		pinger, err := probing.NewPinger("192.168.2.1")
		if err != nil {
			fmt.Println(err.Error())
		}
		pinger.Count = 1
		pinger.Run()
		stats := pinger.Statistics()

		fmt.Println(stats.MaxRtt * time.Millisecond)
		time.Sleep(1 * time.Second)
	}
}
