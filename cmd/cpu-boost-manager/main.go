package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"linux/pkg/battery"
	"linux/pkg/boost"
	"linux/pkg/logger"
	"time"
)

type Battery interface {
	IsCharging() (bool, error)
}

type Booster interface {
	SetStatus(bool) error
	Status() (bool, error)
}

type Policy string

const (
	Auto      Policy = "auto"
	AlwaysOn  Policy = "always-on"
	AlwaysOff Policy = "always-off"
)

type Config struct {
	Policy
}

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logger.Fatal(err.Error())
	}

	viper.SetConfigFile(filepath.Join(homedir, ".config", "uq.conf"))
	viper.SetConfigType("json")

	if err = viper.ReadInConfig(); err != nil {
		logger.Fatal(err.Error())
	}

	config := viper.Get("cpu-boost-manager").(map[string]interface{})

	booster, err := boost.GetBooster()
	if err != nil {
		logger.Fatal(err.Error())
	}

	battery := battery.NewPrimaryBattery()

	if config["policy"] == AlwaysOn {
		booster.SetStatus(true)
	} else if config["policy"] == AlwaysOff {
		booster.SetStatus(false)
	} else {
		for {
			if battery.IsCharging() && booster.Status() {
				err := booster.SetStatus(false)
				if err != nil {
					logger.Warn("can't switch boost mode: %s", err.Error())
				}
			} else if !battery.IsCharging() && !booster.Status() {
				booster.SetStatus(true)
				if err != nil {
					logger.Warn("can't switch boost mode: %s", err.Error())
				}
			}

			time.Sleep(5 * time.Second)
		}
	}
}
