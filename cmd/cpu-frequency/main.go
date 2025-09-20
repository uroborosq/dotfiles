package main

import (
	"fmt"

	"github.com/bitfield/script"
	"github.com/samber/lo"
)

func main() {
	s, _ := script.ListFiles("/sys/devices/system/cpu/cpu*/cpufreq/scaling_cur_freq").
		Concat().
		FilterLine(func(s string) string { return s[:len(s)-3] }).
		Slice()
	fmt.Println(lo.MaxBy(s, func(a, b string) bool { return len(a) > len(b) || (a > b && len(a) == len(b)) }), "MHz")
}
