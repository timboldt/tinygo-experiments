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

// This is an experiment designed for the Adafruit Feather M0, although it could
// be easily adapted if certain machine-specific pins like the battery voltage
// pin were changed.
package main

import (
	"fmt"
	"machine"

	"time"

	"github.com/timboldt/tinygo-experiments/pkg/ina219"

	"tinygo.org/x/drivers/ds3231"
	"tinygo.org/x/drivers/hd44780i2c"
)

func main() {
	//
	// === Initialize hardware ===
	//

	// Initialize common hardware.
	machine.InitADC()
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_100KHZ,
	})

	// Initialize clock.
	clock := ds3231.New(machine.I2C0)
	clock.Configure()

	lcd := hd44780i2c.New(machine.I2C0, 0x27)

	lcd.Configure(hd44780i2c.Config{
		Width:       20, // required
		Height:      4,  // required
		CursorOn:    false,
		CursorBlink: false,
	})

	// Initialize Battery Voltage ADC.
	vbat := machine.ADC{Pin: machine.D9}
	vbat.Configure(machine.ADCConfig{})

	powerMeter := ina219.New(machine.I2C0)
	powerMeter.Configure(nil)

	for {
		//
		// === Read the sensors ===
		//
		temperatureMilliC, err := clock.ReadTemperature()
		if err != nil {
			println("Error reading temperature", err)
		}
		clockTime, err := clock.ReadTime()
		if err != nil {
			println("Error reading time", err)
			clockTime = time.Time{}
		}
		// Voltage divider is half of 3.3V and total scale is 65536.
		batteryMilliVolts := int32(vbat.Get()) * 2 * 3300 / 65536

		milliVolts := powerMeter.GetVoltage_mV()
		microAmps := powerMeter.GetCurrent_uA()
		microWatts := powerMeter.GetPower_uW()

		//
		// === Update the display ===
		//
		statusInfo := fmt.Sprintf("%d.%02dC  %d.%02dV",
			temperatureMilliC/1000, temperatureMilliC%1000/10,
			batteryMilliVolts/1000, batteryMilliVolts%1000/10)
		println(statusInfo)

		powerInfo := fmt.Sprintf("%d.%02dmV %dmA %dmW",
			milliVolts/1000, milliVolts%1000/10,
			microAmps/1000,
			microWatts/1000)
		println(powerInfo)

		clockInfo := clockTime.Format("Mon Jan _2 15:04:05")
		println(clockInfo)

		lcd.ClearDisplay()
		lcd.Print([]byte(statusInfo))
		lcd.Print([]byte("\n"))
		lcd.Print([]byte(powerInfo))
		lcd.Print([]byte("\n"))
		lcd.Print([]byte(clockInfo))

		time.Sleep(1 * time.Second)
	}
}
