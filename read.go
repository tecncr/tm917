package tm917

import (
	"fmt"
	"strings"
)

// Unit represents the temperature unit for the reading.
type Unit string

const (
	UnitFahrenheit Unit = "F"
	UnitCelsius    Unit = "C"
)

const (
	minDataLength = 14
	twoDecimalLen = 5
	oneDecimalLen = 4
)

// Read parses the raw 16-byte data from the thermometer, automatically detecting
// whether it has 1-decimal or 2-decimal precision. It returns the temperature,
// the unit (F or C), the entire raw string, and any error encountered.
func (t *TM917) Read() (float32, Unit, string, error) {
	raw, err := t.Raw()
	if err != nil {
		return 0, "", "", fmt.Errorf("failed to read raw data: %w", err)
	}

	// Strip start/end markers (often \x02 and \r).
	str := strings.Trim(raw, "\x02\r")
	if len(str) < minDataLength {
		return 0, "", str, fmt.Errorf("data string too short: got %d, want >= %d", len(str), minDataLength)
	}

	// Determine the unit (4th character).
	var unit Unit
	switch str[3] {
	case '2':
		unit = UnitFahrenheit
	case '1':
		unit = UnitCelsius
	default:
		return 0, "", str, fmt.Errorf("unknown unit in data string: %q", str)
	}

	// Check precision by the 6th character: '2' => 2 decimals, '1' => 1 decimal.
	var reading float32
	switch str[5] {
	case '2': // 2 decimals
		r, err := parseTemperature(str[9:14], twoDecimalLen)
		if err != nil {
			return 0, unit, str, err
		}
		reading = r
	case '1': // 1 decimal
		r, err := parseTemperature(str[10:14], oneDecimalLen)
		if err != nil {
			return 0, unit, str, err
		}
		reading = r
	default:
		return 0, unit, str, fmt.Errorf("unknown precision flag in data string: %q", str)
	}

	return reading, unit, str, nil
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
	case oneDecimalLen: // 4 => 1 decimal
		// Example: "0826" => 826 => 82.6
		reading = float32(i) / 10
	case twoDecimalLen: // 5 => 2 decimals
		// Example: "08276" => 8276 => 82.76
		reading = float32(i) / 100
	default:
		return 0, fmt.Errorf("unsupported precision length: %d", expectedLen)
	}
	return reading, nil
}
