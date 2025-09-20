package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"text/template"

	"github.com/bitfield/script"
	"github.com/h2non/bimg"
)

func blur(file string) error {
	buf, err := bimg.Read(file)
	if err != nil {
		return err
	}
	img := bimg.NewImage(buf)
	outBuf, err := img.Process(bimg.Options{GaussianBlur: bimg.GaussianBlur{Sigma: 8}, Speed: 8})
	if err != nil {
		return err
	}
	return bimg.Write(file, outBuf)
}

func main() {
	names, err := script.Exec("swaymsg -r -t get_outputs").JQ(`.[] | select(.type == "output" and .active == true) | .name`).Slice()
	if err != nil {
		panic(err)
	}

	bimg.VipsVectorSetEnabled(true)

	_, err = script.Slice(names).ExecForEach("grim -t jpeg -q 80 -o {{ . }} /tmp/swaylock-{{ . }}.png").Slice()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, name := range names {
		wg.Go(func() {
			name := strings.Trim(name, `"`)
			err := blur(fmt.Sprintf("/tmp/swaylock-%s.png", name))
			if err != nil {
				panic(err)
			}
		})
	}
	wg.Wait()

	t, err := template.New("").Parse("swaylock {{ range . }} -i {{ . }}:/tmp/swaylock-{{ . }}.png {{ end }}")
	if err != nil {
		panic(err)
	}

	b := bytes.NewBuffer(nil)
	if err := t.Execute(b, names); err != nil {
		panic(err)
	}

	args, _ := script.Args().Join().String()

	err = script.Exec(b.String() + args).Wait()
	if err != nil {
		panic(err)
	}
}
