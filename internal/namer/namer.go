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

// Package namer provides a namer strategy interface for naming entities.
package namer

import (
	"fmt"

	"github.com/openconfig/entity-naming/oc"
)

// Namer is a strategy interface for naming entities.
type Namer interface {
	// LoopbackInterface returns the name of the loopback interface with the
	// specified zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	LoopbackInterface(index int) (string, error)

	// AggregateInterface returns the name of the aggregate interface with
	// the specified zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	AggregateInterface(index int) (string, error)

	// AggregateMemberInterface returns the name of the member interface
	// bound to the aggregate interface with the specified zero-based index,
	// or an error if no such name exists.
	// This method will never be called with a negative index.
	AggregateMemberInterface(index int) (string, error)

	// Linecard returns the name of the linecard component with the specified
	// zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	Linecard(index int) (string, error)

	// ControllerCard returns the name of the controller card component with the
	// specified zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	ControllerCard(index int) (string, error)

	// Fabric returns the name of the fabric component with the specified
	// zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	Fabric(index int) (string, error)

	// Port returns the name of a physical port with the specified parameters,
	// or an error if no such name exists. This method will never be called with
	// negative index parameters or an unset or unknown port speed.
	Port(port *PortParams) (string, error)

	// Return whether the device has a fixed form factor.
	IsFixedFormFactor() bool
}

// PortParams are parameters of a network port.
type PortParams struct {
	// SlotIndex is the zero-based index of the slot on the device.
	// This value is nil on fixed form factor devices.
	SlotIndex *int
	// PICIndex is the zero-based index of the PIC within the slot.
	PICIndex int
	// PortIndex is the zero-based index of the port within the PIC.
	PortIndex int
	// ChannelIndex is the zero-based index of the channel within the Port.
	// This value is nil for unchannelized ports.
	ChannelIndex *int
	// Speed is the ethernet link speed of the port.
	Speed oc.E_IfEthernet_ETHERNET_SPEED
}

func (pp *PortParams) String() string {
	return fmt.Sprintf("%+v", *pp)
}
