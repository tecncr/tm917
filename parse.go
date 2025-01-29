package tm917

import "fmt"

// parseUnit converts a byte to a Unit (UnitFahrenheit or UnitCelsius).
func parseUnit(b byte) (Unit, error) {
	switch b {
	case '2':
		return UnitFahrenheit, nil
	case '1':
		return UnitCelsius, nil
	default:
		return "", fmt.Errorf("invalid unit: %c", b)
	}
}

// parseReading extracts the precision and temperature from the raw string.
func parseReading(str string) (Precision, float32, error) {
	switch str[precisionPos] {
	case '2':
		temp, err := parseTemperature(str[temp2Start:temp2Start+twoDecimalLen], twoDecimalLen)
		return Precision2Decimal, temp, err
	case '1':
		temp, err := parseTemperature(str[temp1Start:temp1Start+oneDecimalLen], oneDecimalLen)
		return Precision1Decimal, temp, err
	default:
		return 0, 0, fmt.Errorf("invalid precision: %c", str[precisionPos])
	}
}

// parseTemperature converts a substring to a float32 temperature.
func parseTemperature(substr string, length int) (float32, error) {
	var value int
	if _, err := fmt.Sscanf(substr, "%d", &value); err != nil {
		return 0, fmt.Errorf("parse temperature %q: %w", substr, err)
	}
	switch length {
	case twoDecimalLen:
		return float32(value) / 100, nil
	case oneDecimalLen:
		return float32(value) / 10, nil
	default:
		return 0, fmt.Errorf("unsupported precision length: %d", length)
	}
}
