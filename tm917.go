package tm917

import (
	"go.bug.st/serial"
)

// TM917 represents the Lutron TM-917 thermometer
// connected to a serial port, with high precision enabled or disabled.
type TM917 struct {
	// Port is the serial port connected to the thermometer.
	Port serial.Port
	// HighPrecision is true if the thermometer is operating with 2 decimal places
	// or false if it is operating with 1 decimal place.
	HighPrecision bool
}

// Creates a new TM917 connected to the specified serial port,
// with high precision enabled or disabled.
func NewTM917(port serial.Port, highPrecision bool) *TM917 {
	return &TM917{
		Port:          port,
		HighPrecision: highPrecision,
	}
}

// Closes the serial port.
func (t *TM917) Stop() error {
	return t.Port.Close()
}
