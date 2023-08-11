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

// Package juniper provides a Juniper naming implementation.
package juniper

import (
	"fmt"
	"strings"

	"github.com/openconfig/entity-naming/internal/namer"
)

var _ namer.Namer = (*Namer)(nil)

// Namer is a Juniper implementation of the Namer interface.
type Namer struct {
	HardwareModel string
}

// LoopbackInterface is an implementation of namer.LoopbackInterface.
func (n *Namer) LoopbackInterface(index uint) (string, error) {
	if index != 0 {
		return "", fmt.Errorf("Juniper only supports loopback interface zero")
	}
	return "lo0", nil
}

// AggregateInterface is an implementation of namer.AggregateInterface.
func (n *Namer) AggregateInterface(index uint) (string, error) {
	const maxIndex = 1151
	if index > maxIndex {
		return "", fmt.Errorf("Juniper aggregate index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("ae%d", index), nil
}

// AggregateMemberInterface is an implementation of namer.AggregateMemberInterface.
func (n *Namer) AggregateMemberInterface(index uint) (string, error) {
	name, err := n.AggregateInterface(index)
	if err != nil {
		return "", err
	}
	return name + ".0", nil
}

// Linecard is an implementation of namer.Linecard.
func (n *Namer) Linecard(index uint) (string, error) {
	const maxIndex = 7
	if index > maxIndex {
		return "", fmt.Errorf("Juniper linecard index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("FPC%d", index), nil
}

// ControllerCard is an implementation of namer.ControllerCard.
func (n *Namer) ControllerCard(index uint) (string, error) {
	const maxIndex = 1
	if index > maxIndex {
		return "", fmt.Errorf("Juniper controller card index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("RE%d", index), nil
}

// Fabric is an implementation of namer.Fabric.
func (n *Namer) Fabric(index uint) (string, error) {
	const maxIndex = 5
	if index > maxIndex {
		return "", fmt.Errorf("Juniper fabric index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("SIB%d", index), nil
}

// Port is a Juniper implementation of namer.Port.
func (n *Namer) Port(pp *namer.PortParams) (string, error) {
	if !pp.Channelizable {
		return "", fmt.Errorf("Juniper does not support unchannelizable ports")
	}

	var nameBuilder strings.Builder
	nameBuilder.WriteString("et-")
	if pp.SlotIndex == nil {
		nameBuilder.WriteString("0/")
		nameBuilder.WriteString(fmt.Sprintf("%d", pp.PICIndex))
	} else {
		nameBuilder.WriteString(fmt.Sprintf("%d", *pp.SlotIndex))
		nameBuilder.WriteString("/0")
	}
	nameBuilder.WriteString(fmt.Sprintf("/%d", pp.PortIndex))
	if pp.ChannelIndex != nil {
		nameBuilder.WriteString(fmt.Sprintf(":%d", *pp.ChannelIndex))
	}
	return nameBuilder.String(), nil
}

// IsFixedFormFactor is a Juniper implementation of namer.IsFixedFormFactor.
func (n *Namer) IsFixedFormFactor() bool {
	// TODO(juniper): Fill in this implementation.
	return false
}

// QoSForwardingGroups is an implementation of namer.QoSForwardingGroups.
func (n *Namer) QoSForwardingGroups(*namer.QoSParams) (*namer.QoSForwardingGroupNames, error) {
	return &namer.QoSForwardingGroupNames{
		NC1: "7",
		AF4: "6",
		AF3: "4",
		AF2: "3",
		AF1: "2",
		BE1: "0",
		BE0: "1",
	}, nil
}
