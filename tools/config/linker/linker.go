//nolint:wrapcheck
package linker

import (
	"os"
	"os/exec"
)

type SymLinker struct{}

func (l SymLinker) Link(s string, d string) error {
	return os.Symlink(s, d)
}

type HardLinker struct{}

func (l HardLinker) Link(s string, d string) error {
	return os.Link(s, d)
}

type Copier struct{}

func (c Copier) Link(s string, d string) error {
	return exec.Command("cp", s, d).Run()
}
