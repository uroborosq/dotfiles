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
	if status {
		return os.WriteFile(b.path, []byte{'1'}, 0o644)
	}

	return os.WriteFile(b.path, []byte{'0'}, 0o644)
}

func (b *CpufreqBooster) Status() bool {
	output, err := os.ReadFile(b.path)
	if err != nil {
		return false
	}

	return output[0] == '1'
}
