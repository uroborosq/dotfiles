package main

import (
	"fmt"
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


func main() {
	booster, err := boost.GetBooster()
	if err != nil {
		logger.Fatal(err.Error())
	}

	battery := battery.NewPrimaryBattery()

	for {
		fmt.Println(booster.Status(), battery.IsCharging())


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