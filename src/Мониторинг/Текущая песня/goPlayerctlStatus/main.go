package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("playerctl", "metadata", "xesam:title")
	title, err := cmd.Output()

	if err != nil {
		fmt.Println("Ничего не вопроизводится")
		return
	}

	cmd = exec.Command("playerctl", "metadata", "xesam:artist")
	artist, err := cmd.Output()

	if len(artist) == 1 && artist[0] == 10 {
		fmt.Println(string(title))
	} else {
		if len(title) > 30 {
			fmt.Println(string([]rune(string(title))[:30]) + "..." + " | " + string(artist))
		} else {
			fmt.Println(string(title) + " | " + string(artist))
		}
	}

}
