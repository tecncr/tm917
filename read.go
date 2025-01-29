package tm917

import (
	"fmt"
	"strings"
)

// Unit represents the temperature unit for the reading.
type Unit string

// Precision represents the number of decimal places in the temperature reading.
type Precision int

const (
	// UnitFahrenheit represents the Fahrenheit temperature unit.
	UnitFahrenheit Unit = "F"
	// UnitCelsius represents the Celsius temperature unit.
	UnitCelsius Unit = "C"

	// Precision1Decimal represents a precision of 0.1 (1 decimal place).
	Precision1Decimal Precision = 1
	// Precision2Decimal represents a precision of 0.01 (2 decimal places).
	Precision2Decimal Precision = 2
)

const (
	minDataLength = 14
	twoDecimalLen = 5
	oneDecimalLen = 4
)

// Read parses the raw 16-byte data from the thermometer, automatically detecting
// whether it has 1-decimal or 2-decimal precision. It returns the temperature,
// the unit (F or C), the entire raw string, and any error encountered.
func (t *TM917) Read() (float32, Unit, Precision, string, error) {
	raw, err := t.Raw()
	if err != nil {
		return 0, "", 0, "", fmt.Errorf("failed to read raw data: %w", err)
	}

	// Strip start/end markers (often \x02 and \r).
	str := strings.Trim(raw, "\x02\r")
	if len(str) < minDataLength {
		return 0, "", 0, str, fmt.Errorf("data string too short: got %d, want >= %d", len(str), minDataLength)
	}

	// Determine the unit (4th character).
	var unit Unit
	switch str[3] {
	case '2':
		unit = UnitFahrenheit
	case '1':
		unit = UnitCelsius
	default:
		return 0, "", 0, str, fmt.Errorf("unknown unit in data string: %q", str)
	}

	// Check precision (the 6th character)
	var (
		reading   float32
		precision Precision
	)
	switch str[5] {
	case '2':
		precision = Precision2Decimal
		r, err := parseTemperature(str[9:14], twoDecimalLen)
		if err != nil {
			return 0, unit, precision, str, fmt.Errorf("failed to parse temperature: %w", err)
		}
		reading = r
	case '1':
		precision = Precision1Decimal
		r, err := parseTemperature(str[10:14], oneDecimalLen)
		if err != nil {
			return 0, unit, precision, str, fmt.Errorf("failed to parse temperature: %w", err)
		}
		reading = r
	default:
		return 0, unit, 0, str, fmt.Errorf("unknown precision flag in data string: %q", str)
	}

	return reading, unit, precision, str, nil
}

// parseTemperature parses a numeric substring and inserts the decimal
// according to the expected length (4 => 1 decimal, 5 => 2 decimals).
func parseTemperature(substr string, expectedLen int) (float32, error) {
	if len(substr) != expectedLen {
		return 0, fmt.Errorf("invalid substring length: got %d, want %d", len(substr), expectedLen)
	}
	var i int
	if _, err := fmt.Sscanf(substr, "%d", &i); err != nil {
		return 0, fmt.Errorf("failed to parse substring %q: %w", substr, err)
	}

	var reading float32
	switch expectedLen {
	case oneDecimalLen:
		reading = float32(i) / 10
	case twoDecimalLen:
		reading = float32(i) / 100
	default:
		return 0, fmt.Errorf("unsupported precision length: %d", expectedLen)
	}
	return reading, nil
}
