package hwmon_test

import (
	"fmt"
	"linux/pkg/hwmon"
	"testing"
)

func TestLinuxSensorParser_Parse(t *testing.T) {
	parser := hwmon.NewSensorParser()

	sensors, err := parser.Parse()
	if err != nil {
		t.Fail()
	}

	fmt.Println(sensors)
}
