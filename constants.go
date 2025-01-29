package tm917

// Supported temperature units
const (
	// UnitFahrenheit represents the Fahrenheit temperature unit.
	UnitFahrenheit Unit = "F"
	// UnitCelsius represents the Celsius temperature unit.
	UnitCelsius Unit = "C"
)

// Supported precision levels
const (
	// Precision1Decimal represents a precision of 0.1 (1 decimal place).
	Precision1Decimal Precision = 1
	// Precision2Decimal represents a precision of 0.01 (2 decimal places).
	Precision2Decimal Precision = 2
)

// Data format constants
const (
	// Minimum valid data length
	minDataLength = 14
	// Length for 0.01° precision
	twoDecimalLen = 5
	// Length for 0.1° precision
	oneDecimalLen = 4
	// Position of unit indicator
	unitPos = 3
	// Position of precision indicator
	precisionPos = 5
	// Start of temp for 0.1°
	temp1Start = 10
	// Start of temp for 0.01°
	temp2Start = 9
)
