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

var vendorToNamer = map[Vendor]namer.Namer{
	VendorArista:  new(arista.Namer),
	VendorCisco:   new(cisco.Namer),
	VendorJuniper: new(juniper.Namer),
	VendorNokia:   new(nokia.Namer),
}

// DeviceParams are parameters of a network device.
type DeviceParams struct {
	Vendor        Vendor
	HardwareModel string
}

// LoopbackInterface returns the platform-specific name of the loopback
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

// AggregatePort returns the platform-specific name of the aggregate physical
// interface with the given zero-based index.
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

// AggregateInterface returns the platform-specific name of the aggregate
// logical interface with the given zero-based index.
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

func lookupNamer(dp *DeviceParams) (namer.Namer, error) {
	vnamer, ok := vendorToNamer[dp.Vendor]
	if !ok {
		return nil, fmt.Errorf("no Namer for vendor %v", dp.Vendor)
	}
	return vnamer, nil
}
