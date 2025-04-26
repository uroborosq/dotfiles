package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	reader, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}

	file, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}

	defer os.Remove(file.Name())
	defer os.Remove(reader.Name())

	_, err = file.Write(input)
	if err != nil {
		panic(err)
	}

	out, err := exec.Command("kitty", "--class", "terminal-floating", "/usr/bin/bash", "-c", fmt.Sprintf("fzf < %s > %s", file.Name(), reader.Name())).CombinedOutput()
	if err != nil {
		panic(err.Error() + string(out))
	}

	out, err = os.ReadFile(reader.Name())
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(out)
}
