package tm917

import "fmt"

// Parses the raw data. Returns the temperature, the raw string, or an error.
func (t *TM917) Read() (float32, string, error) {
	str, err := t.Raw()
	if err != nil {
		return 0, "", err
	}

	var (
		reading float32
		substr  string
	)

	if t.HighPrecision {
		// Look for "000" then take the next 5 chars.
		idx := findSequence(str, "000")
		if idx < 0 || idx+8 > len(str) {
			return 0, str, fmt.Errorf("invalid high precision reading: %q", str)
		}
		substr = str[idx+3 : idx+8]
	} else {
		// Look for "0000" then take the next 4 chars.
		idx := findSequence(str, "0000")
		if idx < 0 || idx+8 > len(str) {
			return 0, str, fmt.Errorf("invalid reading: %q", str)
		}
		substr = str[idx+4 : idx+8]
	}

	// Insert a decimal point before the last digit of substr.
	if len(substr) == 5 {
		substr = substr[:3] + "." + substr[3:]
	} else {
		return 0, str, fmt.Errorf("malformed substring: %q", substr)
	}

	if _, err = fmt.Sscanf(substr, "%f", &reading); err != nil {
		return 0, str, fmt.Errorf("parse float: %w", err)
	}

	return reading, str, nil
}
