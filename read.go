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

	// Strip leading/trailing whitespace (including \r and \n):
	str := strings.TrimSpace(raw)

	if len(str) < minDataLength {
		return 0, "", str, fmt.Errorf("data string too short: got %d bytes, want >= %d", len(str), minDataLength)
	}

	// Determine unit based on the 4th character ('2' => Fahrenheit, '1' => Celsius).
	var unit Unit
	switch str[3] {
	case '2':
		unit = UnitFahrenheit
	case '1':
		unit = UnitCelsius
	default:
		return 0, "", str, fmt.Errorf("unknown unit in data string: %q", str)
	}

	// Try 2-decimal precision first
	if reading, err := parseTemperature(str[9:14], twoDecimalLen); err == nil {
		return reading, unit, str, nil
	}

	// Fall back to 1-decimal precision
	if reading, err := parseTemperature(str[10:14], oneDecimalLen); err == nil {
		return reading, unit, str, nil
	}

	return 0, unit, str, fmt.Errorf("invalid data string: %q", str)
}

// parseTemperature attempts to parse a temperature string with the given length.
// Returns the parsed temperature or an error if parsing fails.
func parseTemperature(substr string, expectedLen int) (float32, error) {
	if len(substr) != expectedLen {
		return 0, fmt.Errorf("invalid substring length: got %d, want %d", len(substr), expectedLen)
	}

	var reading float32
	insertPos := expectedLen - 2
	formatted := substr[:insertPos] + "." + substr[insertPos:]
	if _, err := fmt.Sscanf(formatted, "%f", &reading); err != nil {
		return 0, fmt.Errorf("failed to parse temperature: %w", err)
	}
	return reading, nil
}
