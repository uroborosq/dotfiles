package main

import (
	"github.com/godbus/dbus/v5"
	"linux/pkg/logger"
	"linux/pkg/monitor"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer conn.Close()

	state, err := monitor.GetCurrentState(conn)
	if err != nil {
		logger.Fatal(err.Error())
	}

	idx, err := monitor.GetBuiltinIndex(state)
	if err != nil {
		logger.Fatal(err.Error())
	}

	switch len(state.LogicalMonitors) {
	case 1:
		err = monitor.TwoMonitors(conn, state, state.Monitors[idx].Info)
		if err != nil {
			logger.Fatal(err.Error())
		}
	case 2:
		err = monitor.ExternalOny(conn, state, state.Monitors[idx].Info)
		if err != nil {
			logger.Fatal(err.Error())
		}
	default:
		logger.Fatal("Unsupported monitor configuration")
	}
}
