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

// Package arista provides an Arista naming implementation.
package arista

import (
	"fmt"
	"strings"

	"github.com/openconfig/entity-naming/internal/namer"
)

var _ namer.Namer = (*Namer)(nil)

// Namer is an Arista implementation of the Namer interface.
type Namer struct {
	HardwareModel string
}

// LoopbackInterface is an implementation of namer.LoopbackInterface.
func (n *Namer) LoopbackInterface(index uint) (string, error) {
	const maxIndex = 1000
	if index > maxIndex {
		return "", fmt.Errorf("Arista loopback index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("Loopback%d", index), nil
}

// AggregateInterface is an implementation of namer.AggregateInterface.
func (n *Namer) AggregateInterface(index uint) (string, error) {
	const maxIndex = 999998
	if index > maxIndex {
		return "", fmt.Errorf("Arista aggregate index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("Port-Channel%d", index+1), nil
}

// AggregateMemberInterface is an implementation of namer.AggregateMemberInterface.
func (n *Namer) AggregateMemberInterface(index uint) (string, error) {
	return n.AggregateInterface(index)
}

// Linecard is an implementation of namer.Linecard.
func (n *Namer) Linecard(index uint) (string, error) {
	const maxIndex = 7
	if index > maxIndex {
		return "", fmt.Errorf("Arista linecard index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("Linecard%d", index+3), nil
}

// ControllerCard is an implementation of namer.ControllerCard.
func (n *Namer) ControllerCard(index uint) (string, error) {
	const maxIndex = 1
	if index > maxIndex {
		return "", fmt.Errorf("Arista controller card index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("Supervisor%d", index+1), nil
}

// Fabric is an implementation of namer.Fabric.
func (n *Namer) Fabric(index uint) (string, error) {
	const maxIndex = 5
	if index > maxIndex {
		return "", fmt.Errorf("Arista fabric index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("Fabric%d", index+1), nil
}

// Port is an implementation of namer.Port.
func (n *Namer) Port(pp *namer.PortParams) (string, error) {
	var nameBuilder strings.Builder
	nameBuilder.WriteString("Ethernet")
	if pp.SlotIndex != nil {
		nameBuilder.WriteString(fmt.Sprintf("%d/", (*pp.SlotIndex)+3))
	}
	nameBuilder.WriteString(fmt.Sprintf("%d", pp.PortIndex))
	if pp.Channelizable {
		if pp.ChannelIndex == nil {
			nameBuilder.WriteString("/1")
		} else {
			nameBuilder.WriteString(fmt.Sprintf("/%d", *pp.ChannelIndex))
		}
	}
	return nameBuilder.String(), nil
}

// IsFixedFormFactor is an implementation of namer.IsFixedFormFactor.
func (n *Namer) IsFixedFormFactor() bool {
	// TODO(arista): Fill in this implementation.
	return false
}

// CommonQoSQueues is an implementation of namer.CommonQoSQueues.
func (n *Namer) CommonQoSQueues(*namer.QoSParams) (*namer.CommonQoSQueueNames, error) {
	return &namer.CommonQoSQueueNames{
		NC1: "NC1",
		AF4: "AF4",
		AF3: "AF3",
		AF2: "AF2",
		AF1: "AF1",
		BE1: "BE1",
		BE0: "BE0",
	}, nil
}
