package tm917

import (
	"fmt"
	"strings"
)

// Read parses the raw 16-byte data from the thermometer, automatically detecting
// whether it has 1-decimal or 2-decimal precision. It returns the temperature,
// the unit (F or C), the entire raw string, and any error encountered.
func (t *TM917) Read() (float32, Unit, Precision, string, error) {
	raw, err := t.Raw()
	if err != nil {
		return 0, "", 0, "", fmt.Errorf("read raw data: %w", err)
	}

	// Clean and validate input
	str := strings.Trim(raw, "\x02\r")
	if len(str) < minDataLength {
		return 0, "", 0, str, fmt.Errorf("data too short: got %d, want %d+", len(str), minDataLength)
	}

	// Parse unit
	unit, err := parseUnit(str[unitPos])
	if err != nil {
		return 0, "", 0, str, err
	}

	// Parse precision and temperature
	precision, temp, err := parseReading(str)
	if err != nil {
		return 0, unit, precision, str, err
	}

	return temp, unit, precision, str, nil
}
