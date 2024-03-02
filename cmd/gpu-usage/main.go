package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	output, err := exec.Command("nvidia-smi", "--query-gpu", "utilization.gpu", "--format", "csv,noheader,nounits").CombinedOutput()
	if err != nil {
		fmt.Printf("error: %s", string(output))
		os.Exit(1)
	}
	fmt.Printf("%s%%\n", strings.TrimSpace(string(output)))
}
