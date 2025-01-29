package tm917

import (
	"fmt"
)

const (
	minDataLength = 14
	twoDecimalLen = 5
	oneDecimalLen = 4
)

// Read parses the raw 16-byte data from the thermometer, automatically detecting
// whether it has 1-decimal or 2-decimal precision. It returns the temperature,
// the entire raw string, and any error encountered.
func (t *TM917) Read() (float32, string, error) {
	str, err := t.Raw()
	if err != nil {
		return 0, "", fmt.Errorf("failed to read raw data: %w", err)
	}

	if len(str) < minDataLength {
		return 0, str, fmt.Errorf("data string too short: got %d bytes, want >= %d", len(str), minDataLength)
	}

	// Try 2-decimal precision first
	if reading, err := parseTemperature(str[9:14], twoDecimalLen); err == nil {
		return reading, str, nil
	}

	// Fall back to 1-decimal precision
	if reading, err := parseTemperature(str[10:14], oneDecimalLen); err == nil {
		return reading, str, nil
	}

	return 0, str, fmt.Errorf("invalid data string: %q", str)
}

// parseTemperature attempts to parse a temperature string with the given length
// Returns the parsed temperature or an error if parsing fails
func parseTemperature(substr string, expectedLen int) (float32, error) {
	if len(substr) != expectedLen {
		return 0, fmt.Errorf("invalid substring length: got %d, want %d", len(substr), expectedLen)
	}

	var reading float32
	insertPos := expectedLen - 2
	formatted := substr[:insertPos] + "." + substr[insertPos:]

	_, err := fmt.Sscanf(formatted, "%f", &reading)
	if err != nil {
		return 0, fmt.Errorf("failed to parse temperature: %w", err)
	}

	return reading, nil
}
