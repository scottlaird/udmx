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

// Package udmx provides a minimal insterface for sending DMX messages
// via a cheap uDMX USB interface.
//
// Supported devices include https://www.amazon.com/dp/B07GT3S6V6 and
// similar hardware.
//
// Inspired by https://github.com/markusb/uDMX-linux
package udmx

import (
	"errors"
	"github.com/google/gousb"
)

// Device describes a USB-connected uDMX device.
type Device struct {
	Dev *gousb.Device
}

// NewUDMXDevice creates a new uDMX device.
func NewUDMXDevice(ctx *gousb.Context) (*Device, error) {

	dev, err := ctx.OpenDeviceWithVIDPID(0x16c0, 0x5dc)
	if err != nil {
		return nil, err
	}

	if m, _ := dev.Manufacturer(); m != "www.anyma.ch" {
		return nil, errors.New("found possible uDMX device but wrong manufacturer string found")
	}

	if dev == nil {
		return nil, errors.New("could not find uDMX device")
	}

	if err != nil {
		return nil, err
	}
	
	return &Device{
		Dev: dev,
	}, nil
}

// Set sends a DMX message to `address` with a value of `value`.
// Generally, this sets a light's brightness, etc to the specified
// value.
//
// Address is a 1-based user-visible DMX address, even though the DMX
// wire format is 0-based.  To set device #5 to 7, call SetDMX(5, 7)
func (d *Device) Set(address, value uint16) error {
	_, err := d.Dev.Control(gousb.ControlOut|gousb.ControlVendor|gousb.ControlDevice, 1, value, address-1, []byte{})
	return err
}
