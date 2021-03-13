// Package ina219 is a driver for the INA219 high side I2C power monitor.
//
// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This is a TinyGo rewrite of the Adafruit library found at https://github.com/adafruit/Adafruit_INA219.
// For completeness, here is the Adafruit attributution and license text for that library:

/* === BEGIN Adafruit attribution text ===
 *!
 * @file Adafruit_INA219.h
 *
 * This is a library for the Adafruit INA219 breakout board
 * ----> https://www.adafruit.com/product/904
 *
 * Adafruit invests time and resources providing this open source code,
 * please support Adafruit and open-source hardware by purchasing
 * products from Adafruit!
 *
 * Written by Bryan Siepert and Kevin "KTOWN" Townsend for Adafruit Industries.
 *
 * BSD license, all text here must be included in any redistribution.
 *
=== END Adafruit attribution text === */

/* === BEGIN Adafruit License text ===
Software License Agreement (BSD License)

Copyright (c) 2012, Adafruit Industries
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
1. Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright
notice, this list of conditions and the following disclaimer in the
documentation and/or other materials provided with the distribution.
3. Neither the name of the copyright holders nor the
names of its contributors may be used to endorse or promote products
derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS ''AS IS'' AND ANY
EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
=== END Adafruit License text === */

package ina219

import (
	"tinygo.org/x/drivers"
)

// Device wraps an I2C connection to a INA219 device.
type Device struct {
	bus     drivers.I2C
	address uint16
	config  Config
}

// Useful configurations are:
//
// 32V 2A (default)
//   CalibrationValue:   4096,
//   ConfigValue:        0x399F,
//   CurrentDivider_mA:  10,
//   PowerMultiplier_uW: 2000,
//
// 16V 400mA (best precision)
//   CalibrationValue:   8192,
//   ConfigValue:        0x019F,
//   CurrentDivider_mA:  20,
//   PowerMultiplier_uW: 1000,

type Config struct {
	CalibrationValue   uint16
	ConfigValue        uint16
	CurrentDivider_mA  uint16
	PowerMultiplier_uW uint16
}

// New creates a new INA219 connection. The I2C bus must already be
// configured.
//
// This function only creates the Device object. It does not touch the device.
func New(bus drivers.I2C) Device {
	return Device{
		bus:     bus,
		address: Address,
		config: Config{
			// Default setup is 32V 2A.
			CalibrationValue:   4096,
			ConfigValue:        0x399F,
			CurrentDivider_mA:  10,
			PowerMultiplier_uW: 2000,
		},
	}
}

// Configure sets up the device for communication.
func (d *Device) Configure(config *Config) bool {
	if config != nil {
		d.config = *config
	}
	d.writeReg(REG_CONFIG, d.config.ConfigValue)
	d.writeReg(REG_CALIBRATION, uint16(d.config.CalibrationValue))
	return true
}

func (d *Device) GetVoltage_mV() int16 {
	val := int16(d.readReg(REG_BUS_VOLTAGE))
	return (val >> 3) * 4
}

func (d *Device) GetCurrent_uA() int32 {
	val := int16(d.readReg(REG_CURRENT))
	return int32(val) * 1000 / int32(d.config.CurrentDivider_mA)
}

func (d *Device) GetPower_uW() int32 {
	val := d.readReg(REG_POWER)
	return int32(val) * int32(d.config.PowerMultiplier_uW)
}

func (d *Device) readReg(reg uint8) uint16 {
	var data [2]byte
	err := d.bus.ReadRegister(uint8(d.address), reg, data[:])
	if err != nil {
		println(err)
		return 0
	}
	return (uint16(data[0]) << 8) + uint16(data[1])
}

func (d *Device) writeReg(reg uint8, val uint16) {
	var data [2]byte
	data[0] = byte(val >> 8)
	data[1] = byte(val & 0xFF)
	err := d.bus.WriteRegister(uint8(d.address), reg, data[:])
	if err != nil {
		println(err)
	}
}
