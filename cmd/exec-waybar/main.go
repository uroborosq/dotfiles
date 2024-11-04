// Copyright (c) 2024, KNS Group LLC ("YADRO").
// All Rights Reserved.
// This software contains the intellectual property of YADRO
// or is licensed to YADRO from third parties. Use of this
// software and the intellectual property contained therein is expressly
// limited to the terms and conditions of the License Agreement under which
// it is provided by YADRO.
package main

import (
	"context"
	"flag"
	"fmt"
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

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	for range limit {
		select {
		case <-ctx.Done():
			log.Infof("Received interrupt signal! Aborting...")
			os.Exit(1)
		default:
			if err := exec.Command("waybar", fmt.Sprintf("--config %s", *path)).Run(); err != nil {
				logger.Warnf(err.Error())
			}
		}
	}
	logger.Fatalf("Exceeded number of attempts: %d", limit)
	os.Exit(1)
}
