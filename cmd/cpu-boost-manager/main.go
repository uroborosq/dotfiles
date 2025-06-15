package main

import (
	"time"

	"github.com/spf13/viper"

	"linux/pkg/battery"
	"linux/pkg/boost"
	"linux/pkg/logger"
)

type Battery interface {
	IsCharging() (bool, error)
}

type Booster interface {
	SetStatus(bool) error
	Status() (bool, error)
}

const (
	Auto      = "auto"
	AlwaysOn  = "always-on"
	AlwaysOff = "always-off"
)

const (
	updateTimeout = 5 * time.Second
)

type Config struct {
	Boost CPUBoostConfig `json:"boost"`
}

type CPUBoostConfig struct {
	Policy string
}

func main() {
	viper.SetConfigFile("/etc/uq.conf")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf(err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logger.Fatalf(err.Error())
	}

	booster, err := boost.GetBooster()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	logger.Infof("Applying policy %s", config.Boost.Policy)

	battery := battery.NewPrimaryBattery()

	switch config.Boost.Policy {
	case AlwaysOn:
		if setError := booster.SetStatus(true); setError != nil {
			logger.Fatalf("can't turn on the boost due to error: %s", setError.Error())
		}
	case AlwaysOff:
		if setError := booster.SetStatus(false); setError != nil {
			logger.Errorf("can't turn off the boost due to error: %s", setError.Error())
		}
	case Auto:
		for {
			if battery.IsCharging() && booster.Status() {
				if setError := booster.SetStatus(false); setError != nil {
					logger.Warnf("can't switch boost mode: %s", setError.Error())
				}
			} else if !battery.IsCharging() && !booster.Status() {
				if setError := booster.SetStatus(true); setError != nil {
					logger.Warnf("can't switch boost mode: %s", setError.Error())
				}
			}

			time.Sleep(updateTimeout)
		}
	}
}
