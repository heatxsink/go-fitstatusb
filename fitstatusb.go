package fitstatusb

import (
	"fmt"

	"go.bug.st/serial"
)

type FitStatUSB struct {
	portName string
	mode     *serial.Mode
}

func New(portName string, baudRate int) (*FitStatUSB, error) {
	fsu := &FitStatUSB{
		portName: portName,
		mode: &serial.Mode{
			BaudRate: baudRate,
		},
	}
	return fsu, nil
}

func (fsu *FitStatUSB) Write(value string) error {
	var err error
	port, err := serial.Open(fsu.portName, fsu.mode)
	if err != nil {
		return err
	}
	_, err = port.Write([]byte(fmt.Sprintf("%s\n\r", value)))
	if err != nil {
		port.Close()
		return err
	}
	port.Close()
	return nil
}
