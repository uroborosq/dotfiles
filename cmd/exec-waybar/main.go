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
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
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
	log.Infof("Starting EXEC-WAYBAR service")

	if !slices.Contains(allowedDE, os.Getenv("XDG_CURRENT_DESKTOP")) {
		log.Fatalf("Can be used only on %s", strings.Join(allowedDE, ", "))
	}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	for range limit {
		select {
		case <-ctx.Done():
			log.Infof("Received interrupt signal! Aboring...")
			os.Exit(1)
		default:
			homeDir, err := os.UserHomeDir()
			if err != nil {
				log.Fatalf("can't determine user home dir: %s", err.Error())
			}
			configPath := filepath.Join(homeDir, ".config", "waybar", "config.json")
			if err := exec.Command("waybar", fmt.Sprintf("--config %s", configPath)).Run(); err != nil {
				logger.Warnf(err.Error())
			}
		}
	}
	logger.Fatalf("Exceeded number of attempts: %d", limit)
	os.Exit(1)
}
