// Copyright 2023 Google LLC
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

// Package naming provides entity-naming functions.
package naming

import (
	"fmt"

	"github.com/openconfig/entity-naming/internal/arista"
	"github.com/openconfig/entity-naming/internal/cisco"
	"github.com/openconfig/entity-naming/internal/juniper"
	"github.com/openconfig/entity-naming/internal/namer"
	"github.com/openconfig/entity-naming/internal/nokia"
)

// Vendor is an enum of network device suppliers.
type Vendor string

// Vendor enum constants.
const (
	VendorArista  = Vendor("Arista")
	VendorCisco   = Vendor("Cisco")
	VendorJuniper = Vendor("Juniper")
	VendorNokia   = Vendor("Nokia")
)

var namerFactories = map[Vendor]func(string) namer.Namer{
	VendorArista:  func(hwm string) namer.Namer { return &arista.Namer{HardwareModel: hwm} },
	VendorCisco:   func(hwm string) namer.Namer { return &cisco.Namer{HardwareModel: hwm} },
	VendorJuniper: func(hwm string) namer.Namer { return &juniper.Namer{HardwareModel: hwm} },
	VendorNokia:   func(hwm string) namer.Namer { return &nokia.Namer{HardwareModel: hwm} },
}

// DeviceParams are parameters of a network device.
type DeviceParams struct {
	Vendor        Vendor
	HardwareModel string
}

// LoopbackInterface returns the vendor-specific name of the loopback
// interface with the given zero-based index.
func LoopbackInterface(dp *DeviceParams, index int) (string, error) {
	namer, err := lookupNamer(dp)
	if err != nil {
		return "", err
	}
	if index < 0 {
		return "", fmt.Errorf("interface index cannot be negative: %d", index)
	}
	return namer.LoopbackInterface(index)
}

// AggregatePort returns the vendor-specific name of the aggregate ports with
// the given zero-based index.
func AggregatePort(dp *DeviceParams, index int) (string, error) {
	namer, err := lookupNamer(dp)
	if err != nil {
		return "", err
	}
	if index < 0 {
		return "", fmt.Errorf("interface index cannot be negative: %d", index)
	}
	return namer.AggregatePort(index)
}

// AggregateInterface returns the vendor-specific name of the interface bound to
// the aggregate port with the given zero-based index.
func AggregateInterface(dp *DeviceParams, index int) (string, error) {
	namer, err := lookupNamer(dp)
	if err != nil {
		return "", err
	}
	if index < 0 {
		return "", fmt.Errorf("interface index cannot be negative: %d", index)
	}
	return namer.AggregateInterface(index)
}

// Linecard returns the vendor-specific name of the linecard with the given
// zero-based index.
func Linecard(dp *DeviceParams, index int) (string, error) {
	namer, err := lookupNamer(dp)
	if err != nil {
		return "", err
	}
	if index < 0 {
		return "", fmt.Errorf("interface index cannot be negative: %d", index)
	}
	return namer.Linecard(index)
}

// ControllerCard returns the vendor-specific name of the controller card with
// the given zero-based index.
func ControllerCard(dp *DeviceParams, index int) (string, error) {
	namer, err := lookupNamer(dp)
	if err != nil {
		return "", err
	}
	if index < 0 {
		return "", fmt.Errorf("interface index cannot be negative: %d", index)
	}
	return namer.ControllerCard(index)
}

// Fabric returns the vendor-specific name of the fabric with the given
// zero-based index.
func Fabric(dp *DeviceParams, index int) (string, error) {
	namer, err := lookupNamer(dp)
	if err != nil {
		return "", err
	}
	if index < 0 {
		return "", fmt.Errorf("interface index cannot be negative: %d", index)
	}
	return namer.Fabric(index)
}

func lookupNamer(dp *DeviceParams) (namer.Namer, error) {
	nf, ok := namerFactories[dp.Vendor]
	if !ok {
		return nil, fmt.Errorf("no Namer for vendor %v", dp.Vendor)
	}
	return nf(dp.HardwareModel), nil
}
