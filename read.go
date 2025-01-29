package tm917

import "fmt"

// Read parses the raw 16-byte data from the thermometer. It extracts the
// temperature (float32), returns the entire raw string, and any error encountered.
func (t *TM917) Read() (float32, string, error) {
	str, err := t.Raw()
	if err != nil {
		return 0, "", err
	}

	// Ensure we got the expected length.
	if len(str) < 16 {
		return 0, str, fmt.Errorf("invalid data length: got %d, want at least 16", len(str))
	}

	var (
		reading float32
		substr  string
	)

	// HighPrecision means we want 2 decimal places.
	// Offset into the 16-byte string by known positions rather than searching for "000" or "0000".
	// Example: "41020200008276" -> last 5 digits "08276" => "82.76".
	if t.HighPrecision {
		substr = str[9:14] // 5 chars (e.g. "08276")
		if len(substr) != 5 {
			return 0, str, fmt.Errorf("invalid high-precision substring: %q", substr)
		}
		substr = substr[:3] + "." + substr[3:]
	} else {
		// For 1 decimal place: "41010100000282" -> last 4 digits "0282" => "28.2".
		substr = str[10:14] // 4 chars (e.g. "0282")
		if len(substr) != 4 {
			return 0, str, fmt.Errorf("invalid substring: %q", substr)
		}
		substr = substr[:2] + "." + substr[2:]
	}

	if _, err := fmt.Sscanf(substr, "%f", &reading); err != nil {
		return 0, str, fmt.Errorf("parse float: %w", err)
	}
	return reading, str, nil
}
