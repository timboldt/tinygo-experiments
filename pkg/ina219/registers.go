package ina219

// Constants/addresses used for I2C.

// The I2C address which this device listens to by default.
const Address = 0x40

// Register names and addresses.
const (
	REG_CONFIG byte = iota
	REG_SHUNT_VOLTAGE
	REG_BUS_VOLTAGE
	REG_POWER
	REG_CURRENT
	REG_CALIBRATION
)

type Calibration int

const (
	V32A2 Calibration = iota
	V32A1
	V16mA400
)
