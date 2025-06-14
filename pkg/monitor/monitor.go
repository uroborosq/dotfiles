package monitor

import (
	"errors"
	"fmt"
	"slices"

	"github.com/godbus/dbus/v5"
	"github.com/kr/pretty"
)

const (
	dbusDestMutterDisplayConfig   = "org.gnome.Mutter.DisplayConfig"
	dbusPathMutterDisplayConfig   = "/org/gnome/Mutter/DisplayConfig"
	dbusMethodGetCurrentState     = "org.gnome.Mutter.DisplayConfig.GetCurrentState"
	dbusMethodApplyMonitorsConfig = "org.gnome.Mutter.DisplayConfig.ApplyMonitorsConfig"
)

type Info struct {
	Connector string
	Vendor    string
	Product   string
	Serial    string
}

type Mode struct {
	Id                 string
	Width              int32
	Height             int32
	RefreshRate        float64
	Scale              float64
	SupportedScales    []float64
	OptionalProperties map[string]dbus.Variant
}

type Monitor struct {
	Info       Info
	Modes      []Mode
	Properties map[string]dbus.Variant
}

type LogicalMonitorResponse struct {
	X          int32
	Y          int32
	Scale      float64
	Transform  uint32
	IsPrimary  bool
	Monitors   []Info
	Properties map[string]interface{}
}

type InfoRequest struct {
	Connector     string
	MonitorModeID string
	Properties    map[string]dbus.Variant
}

type LogicalMonitorRequest struct {
	X         int32
	Y         int32
	Scale     float64
	Transform uint32
	IsPrimary bool
	Monitors  []InfoRequest
}

type State struct {
	Serial          uint32
	Monitors        []Monitor
	LogicalMonitors []LogicalMonitorResponse
	Properties      map[string]dbus.Variant
}

type ApplyStateRequest struct {
	Serial          uint32
	Method          ApplyConfigType
	LogicalMonitors []LogicalMonitorResponse
	Properties      map[string]dbus.Variant
}

func GetCurrentState(conn *dbus.Conn) (State, error) {
	var state State

	obj := conn.Object(dbusDestMutterDisplayConfig, dbusPathMutterDisplayConfig)
	err := obj.Call(dbusMethodGetCurrentState, 0).Store(&state.Serial, &state.Monitors, &state.LogicalMonitors, &state.Properties)

	return state, err
}

func GetBuiltinIndex(state State) (int, error) {
	for i, mon := range state.Monitors {
		if variant, ok := mon.Properties["is-builtin"]; ok {
			isBuiltin, ok := variant.Value().(bool)
			if !ok {
				continue
			}

			if isBuiltin {
				return i, nil
			}
		}
	}

	return 0, errors.New("can't determine builtin monitor")
}

type ApplyConfigType uint32

const (
	Verify ApplyConfigType = iota
	Temporary
	Persistent
)

func ExternalOny(conn *dbus.Conn, state State, builtinMonitorInfo Info) error {
	if len(state.Monitors) != 2 {
		return errors.New("only dual monitor configuration is supported")
	}

	builtinLogicalMonitorIdx := slices.IndexFunc(state.LogicalMonitors, func(logicalMonitor LogicalMonitorResponse) bool {
		return slices.Contains(logicalMonitor.Monitors, builtinMonitorInfo)
	})
	if builtinLogicalMonitorIdx == -1 {
		return errors.New("given logical configuration doesn't contain builtin monitor")
	}

	builtinPhysicalMonitorIdx := slices.IndexFunc(state.Monitors, func(monitor Monitor) bool {
		return monitor.Info == builtinMonitorInfo
	})
	if builtinPhysicalMonitorIdx == -1 {
		return errors.New("given physical configuration doesn't contain builtin monitor")
	}

	logicalMonitors := make([]LogicalMonitorRequest, 1)
	logicalMonitors[0] = LogicalMonitorRequest{
		X:         0,
		Y:         0,
		Scale:     state.LogicalMonitors[1-builtinLogicalMonitorIdx].Scale,
		Transform: state.LogicalMonitors[1-builtinLogicalMonitorIdx].Transform,
		IsPrimary: true,
		Monitors:  make([]InfoRequest, len(state.LogicalMonitors[1-builtinPhysicalMonitorIdx].Monitors)),
	}

	for i, mon := range state.LogicalMonitors[1-builtinLogicalMonitorIdx].Monitors {
		logicalMonitors[0].Monitors[i] = InfoRequest{
			Connector:     mon.Connector,
			MonitorModeID: state.Monitors[1-builtinPhysicalMonitorIdx].Modes[0].Id,
			Properties:    map[string]dbus.Variant{},
		}
	}

	pretty.Println(logicalMonitors)

	mutter := conn.Object(dbusDestMutterDisplayConfig, dbusPathMutterDisplayConfig)

	return mutter.Call(dbusMethodApplyMonitorsConfig, 0, state.Serial, Persistent, logicalMonitors, state.Properties).Err
}

func TwoMonitors(conn *dbus.Conn, state State, builtinMonitorInfo Info) error {
	if len(state.LogicalMonitors) != 1 || len(state.Monitors) != 2 {
		return errors.New("wrong display configuration")
	}

	if slices.Contains(state.LogicalMonitors[0].Monitors, builtinMonitorInfo) {
		return errors.New("wrong display configuration")
	}

	builtinIdx := slices.IndexFunc(state.Monitors, func(monitor Monitor) bool {
		return monitor.Info == builtinMonitorInfo
	})

	logicalMonitors := make([]LogicalMonitorRequest, 2)

	logicalMonitors[0] = resToReq(state.LogicalMonitors[0], state.Monitors[builtinIdx])
	logicalMonitors[0].Scale = 1.25
	logicalMonitors[0].IsPrimary = false

	especiallySmartValue := int32((1 / logicalMonitors[0].Scale) * float64(state.Monitors[builtinIdx].Modes[0].Width))
	fmt.Println(especiallySmartValue)

	logicalMonitors[1] = LogicalMonitorRequest{
		X:         especiallySmartValue,
		Y:         0,
		Scale:     1,
		Transform: 0,
		IsPrimary: true,
		Monitors: []InfoRequest{
			{
				Connector:     state.Monitors[1-builtinIdx].Info.Connector,
				MonitorModeID: state.Monitors[1-builtinIdx].Modes[0].Id,
				Properties:    state.Monitors[1-builtinIdx].Properties,
			},
		},
	}

	mutter := conn.Object(dbusDestMutterDisplayConfig, dbusPathMutterDisplayConfig)

	return mutter.Call(dbusMethodApplyMonitorsConfig, 0, state.Serial, Persistent, logicalMonitors, state.Properties).Err
}

func resToReq(res LogicalMonitorResponse, monitor Monitor) LogicalMonitorRequest {
	result := LogicalMonitorRequest{
		X:         res.X,
		Y:         res.Y,
		Scale:     res.Scale,
		Transform: res.Transform,
		IsPrimary: res.IsPrimary,
		Monitors: []InfoRequest{
			{
				Connector:     monitor.Info.Connector,
				MonitorModeID: monitor.Modes[0].Id,
				Properties:    monitor.Properties,
			},
		},
	}

	return result
}
