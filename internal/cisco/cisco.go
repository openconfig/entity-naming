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

// Package cisco provides a Cisco naming implementation.
package cisco

import (
	"fmt"
	"strings"

	"github.com/openconfig/entity-naming/internal/namer"
	"github.com/openconfig/entity-naming/oc"
)

var _ namer.Namer = (*Namer)(nil)

// Namer is a Cisco implementation of the Namer interface.
type Namer struct {
	HardwareModel string
}

// LoopbackInterface is an implementation of namer.LoopbackInterface.
func (n *Namer) LoopbackInterface(index uint) (string, error) {
	return fmt.Sprintf("Loopback%d", index), nil
}

// AggregateInterface is an implementation of namer.AggregateInterface.
func (n *Namer) AggregateInterface(index uint) (string, error) {
	const maxIndex = 65534
	if index > maxIndex {
		return "", fmt.Errorf("Cisco aggregate index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("Bundle-Ether%d", index+1), nil
}

// AggregateMemberInterface is an implementation of namer.AggregateMemberInterface.
func (n *Namer) AggregateMemberInterface(index uint) (string, error) {
	return n.AggregateInterface(index)
}

// Linecard is an implementation of namer.Linecard.
func (n *Namer) Linecard(index uint) (string, error) {
	const maxIndex = 7
	if index > maxIndex {
		return "", fmt.Errorf("Cisco linecard index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("0/%d/CPU0", index), nil
}

// ControllerCard is an implementation of namer.ControllerCard.
func (n *Namer) ControllerCard(index uint) (string, error) {
	const maxIndex = 1
	if index > maxIndex {
		return "", fmt.Errorf("Cisco controller card index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("0/RP%d/CPU0", index), nil
}

// Fabric is an implementation of namer.Fabric.
func (n *Namer) Fabric(index uint) (string, error) {
	const maxIndex = 7
	if index > maxIndex {
		return "", fmt.Errorf("Cisco fabric index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("0/FC%d", index), nil
}

var speedStrings = map[oc.E_IfEthernet_ETHERNET_SPEED]string{
	oc.IfEthernet_ETHERNET_SPEED_SPEED_10GB:  "TenGig",
	oc.IfEthernet_ETHERNET_SPEED_SPEED_100GB: "HundredGig",
	oc.IfEthernet_ETHERNET_SPEED_SPEED_400GB: "FourHundredGig",
}

// Port is an implementation of namer.Port.
func (n *Namer) Port(pp *namer.PortParams) (string, error) {
	speed, ok := speedStrings[pp.Speed]
	if !ok {
		return "", fmt.Errorf("no known string for port speed %v", pp.Speed)
	}
	var nameBuilder strings.Builder
	nameBuilder.WriteString(speed + "E0/")
	if pp.SlotIndex == nil {
		nameBuilder.WriteString("0")
	} else {
		nameBuilder.WriteString(fmt.Sprintf("%d", *pp.SlotIndex))
	}
	nameBuilder.WriteString(fmt.Sprintf("/0/%d", pp.PortIndex))
	if pp.ChannelIndex != nil {
		nameBuilder.WriteString(fmt.Sprintf("/%d", *pp.ChannelIndex))
	}
	return nameBuilder.String(), nil
}

// IsFixedFormFactor is an implementation of namer.IsFixedFormFactor.
func (n *Namer) IsFixedFormFactor() bool {
	// TODO(cisco): Fill in this implementation.
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

// ServicePorts is an implementation of namer.ServicePorts.
func (n *Namer) ServicePorts() *namer.ServicePorts {
	return &namer.ServicePorts{
		GNMI:  9339,
		GNOI:  9339,
		GRIBI: 9340,
		P4RT:  9559,
	}
}
