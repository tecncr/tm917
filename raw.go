package tm917

import "fmt"

// Reads the raw 16-byte value from the thermometer, as a string.
func (t *TM917) Raw() (string, error) {
	buf := make([]byte, 16)
	_, err := t.Port.Read(buf)
	if err != nil {
		return "", fmt.Errorf("read from device: %w", err)
	}

	if err = t.Port.ResetInputBuffer(); err != nil {
		return "", fmt.Errorf("reset device buffer: %w", err)
	}
	return string(buf), nil
}
