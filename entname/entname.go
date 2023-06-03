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

// Package entname provides entity-naming functions.
package entname

import (
	"fmt"

	"github.com/openconfig/entity-naming/internal/arista"
	"github.com/openconfig/entity-naming/internal/cisco"
	"github.com/openconfig/entity-naming/internal/juniper"
	"github.com/openconfig/entity-naming/internal/namer"
	"github.com/openconfig/entity-naming/internal/nokia"
	"github.com/openconfig/entity-naming/oc"
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

func (dp *DeviceParams) String() string {
	return fmt.Sprintf("%+v", *dp)
}

// PortParams are parameters of a network port.
//
//go:generate ./oc/generate.sh
type PortParams struct {
	SlotIndex, PICIndex, PortIndex, ChannelIndex int
	Channelized                                  bool
	Speed                                        oc.E_IfEthernet_ETHERNET_SPEED
}

func (pp *PortParams) String() string {
	return fmt.Sprintf("%+v", *pp)
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

// AggregateInterface returns the vendor-specific name of the aggregate
// interface with the given zero-based index.
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

// AggregateMemberInterface returns the vendor-specific name of the member
// interface bound to the aggregate interface with the given zero-based index.
func AggregateMemberInterface(dp *DeviceParams, index int) (string, error) {
	namer, err := lookupNamer(dp)
	if err != nil {
		return "", err
	}
	if index < 0 {
		return "", fmt.Errorf("interface index cannot be negative: %d", index)
	}
	return namer.AggregateMemberInterface(index)
}

// Port returns the vendor-specific name of the physical interface with the
// given port parameters.
func Port(dp *DeviceParams, pp *PortParams) (string, error) {
	namer, err := lookupNamer(dp)
	if err != nil {
		return "", err
	}
	npp, err := namerPortParams(pp, namer.IsFixedFormFactor())
	if err != nil {
		return "", err
	}
	return namer.Port(npp)
}

func namerPortParams(pp *PortParams, fixedFormFactor bool) (*namer.PortParams, error) {
	if pp.SlotIndex < 0 {
		return nil, fmt.Errorf("slot index cannot be negative: %d", pp.SlotIndex)
	}
	if pp.PICIndex < 0 {
		return nil, fmt.Errorf("pic index cannot be negative: %d", pp.PICIndex)
	}
	if pp.PortIndex < 0 {
		return nil, fmt.Errorf("port index cannot be negative: %d", pp.PortIndex)
	}
	if pp.ChannelIndex < 0 {
		return nil, fmt.Errorf("channel index cannot be negative: %d", pp.ChannelIndex)
	}
	if pp.SlotIndex > 0 && fixedFormFactor {
		return nil, fmt.Errorf("cannot have a non-zero slot index on a fixed form factor device")
	}
	if pp.ChannelIndex > 0 && !pp.Channelized {
		return nil, fmt.Errorf("cannot have a non-zero channel index with an unchannelized port")
	}
	if pp.Speed == oc.IfEthernet_ETHERNET_SPEED_UNSET || pp.Speed == oc.IfEthernet_ETHERNET_SPEED_SPEED_UNKNOWN {
		return nil, fmt.Errorf("port speed cannot be unset or unknown")
	}
	npp := &namer.PortParams{
		PICIndex:  pp.PICIndex,
		PortIndex: pp.PortIndex,
		Speed:     pp.Speed,
	}
	if !fixedFormFactor {
		npp.SlotIndex = &pp.SlotIndex
	}
	if pp.Channelized {
		npp.ChannelIndex = &pp.ChannelIndex
	}
	return npp, nil
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
