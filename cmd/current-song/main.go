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
	limit := flag.Int("limit", 100, "limit output")
	flag.Parse()

	cmd := exec.Command("playerctl", "metadata", "xesam:title")
	title, err := cmd.Output()
	if err != nil {
		fmt.Println(" Ничего не вопроизводится")
		return
	}

	cmd = exec.Command("playerctl", "metadata", "xesam:artist")
	artist, _ := cmd.Output()
	status, _ := exec.Command("playerctl", "status").Output()

	artist = bytes.TrimSpace(artist)
	title = bytes.TrimSpace(title)

	var output strings.Builder

	if slices.Equal(status, []byte("Playing\n")) {
		output.WriteString(" ")
	} else if slices.Equal(status, []byte("Paused\n")) {
		output.WriteString(" ")
	}

	if len(title) > *limit {
		output.WriteString(string([]rune(string(title)[:*limit])) + "...")
	} else {
		output.Write(title)
	}

	if len(artist) > *limit {
		output.WriteString(" | ")
		output.WriteString(string([]rune(string(artist)[:*limit])) + "...")
	} else if len(artist) != 0 {
		output.WriteString(" | ")
		output.Write(artist)
	}
	fmt.Println(output.String())
}
