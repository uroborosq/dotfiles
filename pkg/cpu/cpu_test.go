package cpu_test

import (
	"fmt"
	"testing"

	monitor "cpu-monitor"
)


func TestCpuStat (t *testing.T) {
	fmt.Println(monitor.GetValue())
}