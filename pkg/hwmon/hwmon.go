package hwmon

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type LinuxSensorParser struct {
	hwmonPath string
}

func NewSensorParser() *LinuxSensorParser {
	return &LinuxSensorParser{hwmonPath: "/sys/class/hwmon/"}
}

func (p *LinuxSensorParser) Parse() (map[string]map[string]float64, error) {
	monitors, err := os.ReadDir(p.hwmonPath)
	if err != nil {
		return nil, err
	}
	sensors := make(map[string]map[string]float64, len(monitors))
	for _, monitor := range monitors {
		//if !monitor.IsDir() {
		//	continue
		//}
		name, sensor, err := p.parseLinuxMonitor(monitor)
		if err != nil {
			continue
		}
		sensors[name] = sensor

	}
	return sensors, nil
}

type Sensor struct {
	Name  string
	Value float64
}

func (p *LinuxSensorParser) parseLinuxMonitor(dir os.DirEntry) (string, map[string]float64, error) {
	var driverName string
	sensors := make([]Sensor, 10)
	sensorsMap := make(map[string]float64, 10)
	path := p.hwmonPath + dir.Name()
	files, err := os.ReadDir(path)
	if err != nil {
		return "", nil, err
	}
	for _, file := range files {
		name := file.Name()

		if name == "name" {
			content, err := os.ReadFile(filepath.Join(path, name))
			if err != nil {
				continue
			}
			driverName = strings.TrimSpace(string(content))
			continue
		}

		number, typeFile, err := parseSensorName(file.Name())
		if err != nil {
			continue
		}

		if number > len(sensors) {
			sensors = append(sensors, make([]Sensor, number-len(sensors))...)
		}

		content, err := os.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			continue
		}
		switch typeFile {
		case sensorName:
			nameSensor := strings.TrimSpace(string(content))
			sensors[number].Name = nameSensor
			sensorsMap[nameSensor] = sensors[number].Value
		case sensorValue:
			value, err := strconv.ParseFloat(strings.TrimSpace(string(content)), 64)
			if err != nil {
				continue
			}
			sensors[number].Value = value
			if sensors[number].Name == "" {
				continue
			}
			sensorsMap[sensors[number].Name] = value
		}
	}

	if len(sensors) != len(sensorsMap) {
		for i, sensor := range sensors {
			if sensor.Name == "" && sensor.Value != 0 {
				sensorsMap[fmt.Sprintf("Sensor %d", i)] = sensor.Value
			}
		}
	}

	return driverName, sensorsMap, err
}

type fileType int

const (
	sensorName fileType = iota + 1
	sensorValue
)

const temp = "temp"

func parseSensorName(name string) (int, fileType, error) {
	// yes, it's ugly, but instead we get faster algorithm
	for i := range temp {
		if name[i] != temp[i] {
			return 0, 0, errors.New("gop")
		}
	}
	var buffer strings.Builder
	var counter int
	for _, r := range name[4:] {
		counter++
		if r == '_' {
			break
		}
		buffer.WriteRune(r)

	}
	number, err := strconv.Atoi(buffer.String())
	if err != nil {
		return 0, 0, err
	}

	if len(name) < 4+counter {
		return 0, 0, errors.New("lol")
	}

	typeFile := name[4+counter:]
	switch typeFile {
	case "input":
		return number, sensorValue, nil
	case "label":
		return number, sensorName, nil
	default:
		return 0, 0, errors.New("heh")
	}
}
