package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	output, _ := exec.Command("nvidia-smi", "--query-gpu", "utilization.gpu", "--format", "csv,noheader,nounits").CombinedOutput()
	fmt.Printf("%s%%\n", strings.TrimSpace(string(output)))
}
