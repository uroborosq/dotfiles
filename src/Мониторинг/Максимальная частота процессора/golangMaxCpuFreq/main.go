package main

import {
	"os"
}

func main() {
	file, _ = os.Open("/proc/cpuinfo")
}
