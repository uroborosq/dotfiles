package main

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"os/signal"
	"slices"
	"strings"
	"syscall"

	"github.com/charmbracelet/log"

	"linux/pkg/logger"
)

const (
	limit = 1000
)

var allowedDE = []string{"sway", "Hyprland"}

func main() {
	path := flag.String("path", "", "path to waybar config")
	flag.Parse()

	logger.Infof("Starting EXEC-WAYBAR service")

	if !slices.Contains(allowedDE, os.Getenv("XDG_CURRENT_DESKTOP")) {
		log.Fatalf("Can be used only on %s", strings.Join(allowedDE, ", "))
	}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT)

	for range limit {
		select {
		case <-ctx.Done():
			log.Infof("Received interrupt signal! Aborting...")
			os.Exit(1)
		default:
			if err := exec.Command("waybar", "--config "+*path).Run(); err != nil {
				logger.Warnf(err.Error())
			}
		}
	}

	logger.Fatalf("Exceeded number of attempts: %d", limit)
	os.Exit(1)
}
