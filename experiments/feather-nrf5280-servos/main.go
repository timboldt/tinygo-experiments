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

// This is an experiment designed for the Adafruit nrf52840 Express.
package main

import (
	"fmt"
	"machine"

	"time"

	"github.com/timboldt/tinygo-experiments/pkg/pca9685"
)

func main() {
	status := "OK"

	//
	// === Initialize hardware ===
	//
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_100KHZ,
	})

	pwm := pca9685.New(machine.I2C0)
	if err := pwm.Configure(); err != nil {
		status = fmt.Sprintf("configure failed: %v", err)
	}

	var pin byte
	for {
		if err := pwm.SetPin(pin, 1250); err != nil {
			status = fmt.Sprintf("set pin PWM failed: %v", err)
		}
		if pin == 0 {
			pin = 1
		} else {
			pin = 0
		}
		if err := pwm.SetPin(pin, 1500); err != nil {
			status = fmt.Sprintf("set pin PWM failed: %v", err)
		}

		fmt.Printf("Status:  %v\n", status)
		time.Sleep(1 * time.Second)
	}
}
