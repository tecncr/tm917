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
    "fmt"

    "go.bug.st/serial"
    "github.com/tecncr/tm917"
)

func main() {
    // Open serial port (Linux example, COM# for Windows)
    // See https://pkg.go.dev/go.bug.st/serial for more information
    port, err := serial.Open("/dev/ttyUSB0", &serial.Mode{})
    if err != nil {
        panic(err)
    }

    // Create new TM917 instance
    thermometer := tm917.NewTM917(port)
    defer thermometer.Stop()

    // Read temperature
    temp, _, err := thermometer.Read()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Temperature: %.2f°C\n", temp)
}
```

## Features

- Read temperature values from Lutron TM-917 thermometer
- Support for both 1 decimal and 2 decimals readings
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
