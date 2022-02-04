package fitstatusb

import (
	"fmt"

	"go.bug.st/serial"
)

type FitStatUSB struct {
	baudRate int
	portName string
	port     serial.Port
}

func New() (*FitStatUSB, error) {
	fsu := &FitStatUSB{
		baudRate: 9600,
		portName: "/dev/ttyACM0",
	}
	mode := &serial.Mode{
		BaudRate: fsu.baudRate,
	}
	var err error
	fsu.port, err = serial.Open(fsu.portName, mode)
	if err != nil {
		return nil, err
	}
	return fsu, nil
}

func (fsu *FitStatUSB) Write(value string) error {
	_, err := fsu.port.Write([]byte(fmt.Sprintf("%s\n\r", value)))
	if err != nil {
		return err
	}
	return nil
}
