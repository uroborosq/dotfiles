package main

import (
	"fmt"
	"os/exec"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	cmd := exec.Command("playerctl", "metadata", "xesam:title")
	title, err := cmd.Output()

	if err != nil {
		fmt.Println("Ничего не вопроизводится")
		return
	}

	cmd = exec.Command("playerctl", "metadata", "xesam:artist")
	artist, _ := cmd.Output()
	status, _ := exec.Command("playerctl", "status").Output()
	if slices.Equal(status, []byte("Playing\n")) {
		fmt.Print("⏵︎ ")
	} else if slices.Equal(status, []byte("Paused\n")) {
		fmt.Print("⏸︎ ")
	}
	if len(artist) == 1 && artist[0] == 10 {
		fmt.Println(string(title))
	} else {
		limit := 100
		if len(title) > limit {
			fmt.Println(string([]rune(string(title))[:limit]) + "..." + " | " + string(artist))
		} else {
			fmt.Println(strings.TrimSuffix(string(title), "\n") + " | " + strings.TrimSuffix(string(artist), "\n"))
		}
	}

}
