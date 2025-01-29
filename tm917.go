package tm917

import (
	"go.bug.st/serial"
)

// TM917 represents the Lutron TM-917 thermometer connected to a serial port.
type TM917 struct {
	// Port is the serial port where the thermometer is connected.
	// See https://pkg.go.dev/go.bug.st/serial#Port for more information.
	Port serial.Port
}

// Creates a new TM917 connected to the specified serial port.
func NewTM917(port serial.Port) *TM917 {
	return &TM917{
		Port: port,
	}
}

// Closes the serial port.
func (t *TM917) Stop() error {
	return t.Port.Close()
}
