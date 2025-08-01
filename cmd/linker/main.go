package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func syncDir(path string) error {
	dir, err := os.ReadDir(filepath.Join(path, "home"))
	if err != nil {
		return fmt.Errorf("failed to read dir %s: %w", err)
	}

	for _m 
}

func syncTo(path string) error {
	userConfig, err := os.ReadDir(filepath.Join(path, "home"))
	if err != nil {
		return err
	}

	for _, entry := range userConfig {

	}

}

func main() {
	path := flag.String("p", "configs", "path to stored configs")
	action := flag.String("a", "", "action to do")

	flag.Parse()

	switch *action {
	case "to":
	case "from":
	default:
		panic("unknown action")
	}

}
