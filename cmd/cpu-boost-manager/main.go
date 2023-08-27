package main

import (
	"linux/pkg/battery"
	"linux/pkg/boost"
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
	booster := boost.GetBooster()
	battery := battery.NewPrimaryBattery()

	boost := booster.Status()

	for {
		
		if battery.IsCharging() && boost {
			booster.SetStatus(false)
		} else if !battery.IsCharging() && !boost {
			booster.SetStatus(true)
		}

		time.Sleep(5 * time.Second)
	}
}