package boost

import "os"

type Booster interface {
	SetStatus(bool) error
	Status() bool
}

func GetBooster() Booster {
	return nil
}

func isFileExists(path string) Booster {
	if _, err := os.Stat(path); err == nil {
		return NewCpufreqBooster()
	}
	return nil
}