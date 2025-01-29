# tm917

[![Go Reference](https://pkg.go.dev/badge/github.com/tecncr/tm917.svg)](https://pkg.go.dev/github.com/tecncr/tm917)

A Go package for reading data from the Lutron TM-917 precision thermometer via serial port

## Installation

```bash
go get github.com/tecncr/tm917
```

## Example

```go
package main

import (
	"log"
	"time"

	"github.com/tecncr/tm917"
	"go.bug.st/serial"
)

func main() {
    // Open serial port with default settings (Linux example, use COM# on Windows)
	// See https://pkg.go.dev/go.bug.st/serial#Open for more information
	port, err := serial.Open("/dev/ttyUSB0", &serial.Mode{})
	if err != nil {
		log.Fatalf("Failed to open port: %v", err)
	}

	// Create new TM917 instance
	thermometer := tm917.NewTM917(port)
	defer thermometer.Stop()

	// Continuously read temperature values
	// The loop can be stopped with Ctrl+C
	for {
		temp, unit, precision, raw, err := thermometer.Read()
		if err != nil {
			log.Printf("Error reading thermometer: %v\n", err)
		} else {
			// Format output based on device precision setting
			// Precision1Decimal = 0.1°, Precision2Decimal = 0.01°
			switch precision {
			case tm917.Precision1Decimal:
				log.Printf("Temperature: %.1f°%s (raw: %q)\n", temp, unit, raw)
			case tm917.Precision2Decimal:
				log.Printf("Temperature: %.2f°%s (raw: %q)\n", temp, unit, raw)
			}
		}
		// Wait before next reading to avoid flooding the serial port
		time.Sleep(500 * time.Millisecond)
	}
}
```

## Features

- Read temperature values from Lutron TM-917 thermometer
- Automatic detection of unit (°C or °F) and precision (0.1 or 0.01)
- Raw data access for advanced usage
- Simple error handling

## Dependencies

- [go.bug.st/serial](https://pkg.go.dev/go.bug.st/serial) - Serial port communication

## Requirements

- Go 1.23 or later
- Physical access to a serial port
- Lutron TM-917 thermometer connected to the serial port

## Documentation

For detailed documentation, see the [Go Reference](https://pkg.go.dev/github.com/tecncr/tm917).

## License

This project is licensed under the MIT License, see the LICENSE file for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
