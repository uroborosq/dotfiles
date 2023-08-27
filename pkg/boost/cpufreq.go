package boost

import "os"

func NewCpufreqBooster() *CpufreqBooster {
	return &CpufreqBooster{
		path: "/sys/devices/system/cpu/cpufreq/boost",
	}
}

type CpufreqBooster struct {
	path string
}

func (b *CpufreqBooster) SetStatus(status bool) error {
	file, err := os.Open(b.path)
	if err != nil {
		return err
	}
	if status {
	_, err = file.Write([]byte{'1'})
	} else {
		_, err = file.Write([]byte{'0'})
	}
	return err
}

func (b *CpufreqBooster) Status() bool {
	output, err := os.ReadFile(b.path)
	if err != nil {
		return false
	}
	return output[0] == '1'
}