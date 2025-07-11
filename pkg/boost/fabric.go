package boost

import (
	"errors"
	"os"
)

type Booster interface {
	SetStatus(bool) error
	Status() bool
}

func GetBooster() (Booster, error) {
	if isFileExists("/sys/devices/system/cpu/cpufreq/boost") {
		return NewCpufreqBooster(), nil
	}

	return nil, errors.New("can't determine booster type")
}

func isFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}
