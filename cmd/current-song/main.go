package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	limit := flag.Int("limit", 50, "limit output")
	flag.Parse()

	cmd := exec.Command("playerctl", "metadata", "xesam:title")
	titleBytes, err := cmd.Output()
	if err != nil {
		fmt.Println(" Ничего не вопроизводится")
		return
	}

	cmd = exec.Command("playerctl", "metadata", "xesam:artist")
	artistBytes, _ := cmd.Output()
	status, _ := exec.Command("playerctl", "status").Output()

	artist := []rune(string(bytes.TrimSpace(artistBytes)))
	title := []rune(string(bytes.TrimSpace(titleBytes)))

	var output strings.Builder

	if slices.Equal(status, []byte("Playing\n")) {
		output.WriteString(" ")
	} else if slices.Equal(status, []byte("Paused\n")) {
		output.WriteString(" ")
	}

	if len(title) > *limit {
		output.WriteString((string(title[:*limit]) + "..."))
	} else {
		output.WriteString(string(title))
	}
	
	if len(artist) > *limit {
		output.WriteString(" | ")
		output.WriteString(string(artist[:*limit]) + "...")
	} else if len(artist) != 0 {
		output.WriteString(" | ")
		output.WriteString(string(artist))
	}
	fmt.Println(output.String())
}
