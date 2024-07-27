package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
	"unsafe"
)

const (
	delay             = time.Millisecond * 50
	red          byte = 0
	green        byte = 255
	blue         byte = 255
	deviceNumber      = "7"
)

func i2csetPerDevice(deviceNumber string, stickNumber string, first string, second string) error {
	output, err := exec.Command("i2cset", "-y", deviceNumber, stickNumber, first, second).CombinedOutput()
	if err != nil {
		return fmt.Errorf("i2cet -y %s %s %s %s finished with error: %s, %s", deviceNumber, stickNumber, first, second, unsafe.String(&output[0], len(output)), err.Error())
	}
	return nil
}

func setStaticColorToSingleStick(deviceNumber string, stickAddress string, red byte, green byte, blue byte) error {
	err := i2csetPerDevice(deviceNumber, stickAddress, "0x08", "0x53")
	if err != nil {
		return err
	}
	time.Sleep(delay)
	err = i2csetPerDevice(deviceNumber, stickAddress, "0x09", "0x00")
	if err != nil {
		return err
	}
	time.Sleep(delay)
	err = i2csetPerDevice(deviceNumber, stickAddress, "0x31", fmt.Sprintf("0x%x", red))
	if err != nil {
		return err
	}
	time.Sleep(delay)
	err = i2csetPerDevice(deviceNumber, stickAddress, "0x32", fmt.Sprintf("0x%x", green))
	if err != nil {
		return err
	}
	time.Sleep(delay)
	err = i2csetPerDevice(deviceNumber, stickAddress, "0x33", fmt.Sprintf("0x%x", blue))
	if err != nil {
		return err
	}
	time.Sleep(delay)
	err = i2csetPerDevice(deviceNumber, stickAddress, "0x08", "0x44")
	if err != nil {
		return err
	}
	time.Sleep(delay)
	return nil
}

func setOrange() error {
	err := setStaticColorToSingleStick(deviceNumber, "0x61", red, green, blue)
	if err != nil {
		return err
	}
	return setStaticColorToSingleStick(deviceNumber, "0x63", red, green, blue)
}

func turnOff() error {
	err := setStaticColorToSingleStick(deviceNumber, "0x61", 0, 0, 0)
	if err != nil {
		return err
	}
	return setStaticColorToSingleStick(deviceNumber, "0x63", 0, 0, 0)
}

func main() {
	offFlag := flag.Bool("off", false, "turn off rgb lighting")
	onFlag := flag.Bool("on", false, "turn on with 255 60 0")
	flag.Parse()
	if *onFlag && *offFlag {
		fmt.Println("error")
		os.Exit(1)
	} else if *onFlag {
		err := setOrange()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else if *offFlag {
		err := turnOff()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Println("no flags used")
		os.Exit(1)
	}
	fmt.Println("ok")
}
