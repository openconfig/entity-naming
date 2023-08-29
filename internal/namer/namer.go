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
	LoopbackInterface(index uint) (string, error)

	// AggregateInterface returns the name of the aggregate interface with
	// the specified zero-based index, or an error if no such name exists.
	AggregateInterface(index uint) (string, error)

	// AggregateMemberInterface returns the name of the member interface
	// bound to the aggregate interface with the specified zero-based index,
	// or an error if no such name exists.
	AggregateMemberInterface(index uint) (string, error)

	// Linecard returns the name of the linecard component with the specified
	// zero-based index, or an error if no such name exists.
	Linecard(index uint) (string, error)

	// ControllerCard returns the name of the controller card component with the
	// specified zero-based index, or an error if no such name exists.
	ControllerCard(index uint) (string, error)

	// Fabric returns the name of the fabric component with the specified
	// zero-based index, or an error if no such name exists.
	Fabric(index uint) (string, error)

	// Port returns the name of a physical port with the specified parameters,
	// or an error if no such name exists. This method will never be called with
	// an unset or unknown port speed.
	Port(port *PortParams) (string, error)

	// Return whether the device has a fixed form factor.
	IsFixedFormFactor() bool

	// CommonQoSQueues returns the names of the common QoS queues, or an error
	// if no such names exist. See the common QoS queue definitions here:
	// https://github.com/openconfig/entity-naming/blob/main/README.md#common-qos-queues
	CommonQoSQueues(qos *QoSParams) (*CommonQoSQueueNames, error)
}

// PortParams are parameters of a network port.
type PortParams struct {
	// SlotIndex is the zero-based index of the slot on the device.
	// This value is nil on fixed form factor devices.
	SlotIndex *uint
	// PICIndex is the zero-based index of the PIC within the slot.
	PICIndex uint
	// PortIndex is the zero-based index of the port within the PIC.
	PortIndex uint
	// ChannelIndex is the zero-based index of the channel within the Port.
	// This value is nil for unchannelized ports.
	ChannelIndex *uint
	// Channelizable indicates whether the port can be channelized.
	Channelizable bool
	// Speed is the ethernet link speed of the port.
	Speed oc.E_IfEthernet_ETHERNET_SPEED
}

func (pp *PortParams) String() string {
	return fmt.Sprintf("%+v", *pp)
}

// QoSParams are parameters of a QoS configuration.
type QoSParams struct {
	NumStrictPriority, NumWeightedRoundRobin uint
}

// CommonQoSQueueNames are the names of common QoS queues.
type CommonQoSQueueNames struct {
	NC1, AF4, AF3, AF2, AF1, BE1, BE0 string
}

func (qn *CommonQoSQueueNames) String() string {
	return fmt.Sprintf("%+v", *qn)
}
