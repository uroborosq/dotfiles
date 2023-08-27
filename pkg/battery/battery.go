package battery

import (
	"os"
	"strings"
)


type Battery interface {
	IsCharging() bool
}

type PrimaryBattery struct {
	path string
}

func NewPrimaryBattery() PrimaryBattery {
	return PrimaryBattery{
		path: "/sys/class/power_supply/BAT0/status",
	}
}

func (b *PrimaryBattery) IsCharging() bool {
	output, err := os.ReadFile(b.path)
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) != "Discharging"
}
