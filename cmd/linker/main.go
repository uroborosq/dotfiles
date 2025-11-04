package main

import (
	"flag"
	"fmt"

	"github.com/bitfield/script"
)

// func syncDir(path string) error {
// 	dir, err := os.ReadDir(filepath.Join(path, "home"))
// 	if err != nil {
// 		return fmt.Errorf("failed to read dir %s: %w", err)
// 	}
//
// }
//
// func syncTo(path string) error {
// 	userConfig, err := os.ReadDir(filepath.Join(path, "home"))
// 	if err != nil {
// 		return err
// 	}
//
// 	for _, entry := range userConfig {
//
// 	}
//
// }

func main() {
	path := flag.String("p", "configs", "path to stored configs")
	action := flag.String("a", "", "action to do")

	flag.Parse()
	files, err := script.FindFiles(*path).Slice()
	if err != nil {
		panic(err)
	}

	switch *action {
	case "to":
		err = script.Slice(files).Basename().ExecForEach(`mkdir -p {{ . }}`).Wait()
		if err != nil {
			panic(err)
		}

		templ := fmt.Sprintf(`cp -f {{ . }} `)
		err = script.Slice(files).ExecForEach()

	case "from":
	default:
		panic("unknown action")
	}

}
