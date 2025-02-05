package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"linux/tools/config/linker"
)

type Linker interface {
	Link(string, string) error
}

func flags() string {
	path := flag.String("source", "configs", "path to stored configs")

	flag.Parse()

	return *path
}

func syncDir(source string, target string, linker Linker) error {
	var f func(string) error

	f = func(prefix string) error {
		entries, err := os.ReadDir(source)
		if err != nil {
			return fmt.Errorf("can't open dir %s: %w", source, err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				if err := f(filepath.Join(filepath.Join(prefix, entry.Name()))); err != nil {
					return err
				}
			} else {
				src := filepath.Join(source, prefix, entry.Name())
				dst := filepath.Join(target, prefix, entry.Name())

				if err := linker.Link(src, dst); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return f("")
}

func main() {
	sourcePath := flags()
	systemPath := filepath.Join(sourcePath, "root")
	userPath := filepath.Join(sourcePath, "home")

	syncDir(systemPath, "/", linker.SymLinker{})
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	syncDir(userPath, currentUser.HomeDir, linker.SymLinker{})
}
