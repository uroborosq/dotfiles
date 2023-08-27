package main

import (
	"fmt"
	
	"os/exec"
)

func main() {
	cmd := exec.Command("systemctl", "is-enabled", "suspend.target")
	result, _ := cmd.Output()
	status := string(result)
	if status == "static\n" {
		fmt.Println("suspend on")
	} else if status == "masked\n" {
		fmt.Println("suspend off")
	} else {
		fmt.Println("Error")
	}
}