package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Linker interface {
	Link(string, string) error
}

func flags() string {
	path := flag.String("source", "configs", "path to stored configs")

	flag.Parse()

	return *path
}

func syncDir(path string, root string, linker Linker) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("can't open dir %s: %w", path, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if err := syncDir(filepath.Join(path, entry.Name())); err != nil {
				return err
			}
		} else {
			os.Link()
			os.C
		}
	}
}

func main() {
	sourcePath := flags()

	systemWide, err := os.ReadDir(filepath.Join(sourcePath, "root"))
	if err != nil {
		panic(err)
	}

	userWide, err := os.ReadDir(filepath.Join(sourcePath, "home"))
	if err != nil {
		panic(err)
	}
}
