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

type Config struct {
	Boost CpuBoostConfig `json:"boost"`
}

type CpuBoostConfig struct {
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
		booster.SetStatus(true)
	case AlwaysOff:
		booster.SetStatus(false)
	case Auto:
		for {
			if battery.IsCharging() && booster.Status() {
				err := booster.SetStatus(false)
				if err != nil {
					logger.Warnf("can't switch boost mode: %s", err.Error())
				}
			} else if !battery.IsCharging() && !booster.Status() {
				booster.SetStatus(true)
				if err != nil {
					logger.Warnf("can't switch boost mode: %s", err.Error())
				}
			}

			time.Sleep(5 * time.Second)
		}
	}
}
