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

	"github.com/openconfig/entity-naming/internal/namer"
)

var _ namer.Namer = (*Namer)(nil)

// Namer is a Cisco implementation of the Namer interface.
type Namer struct {
	HardwareModel string
}

// LoopbackInterface is a implementation of namer.LoopbackInterface.
func (n *Namer) LoopbackInterface(index int) (string, error) {
	return fmt.Sprintf("Loopback%d", index), nil
}

// AggregateInterface is a implementation of namer.AggregateInterface.
func (n *Namer) AggregateInterface(index int) (string, error) {
	const maxIndex = 65534
	if index > maxIndex {
		return "", fmt.Errorf("Cisco aggregate index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("Bundle-Ether%d", index+1), nil
}

// AggregateMemberInterface is a implementation of namer.AggregateMemberInterface.
func (n *Namer) AggregateMemberInterface(index int) (string, error) {
	return n.AggregateInterface(index)
}

// Linecard is a implementation of namer.Linecard.
func (n *Namer) Linecard(index int) (string, error) {
	const maxIndex = 7
	if index > maxIndex {
		return "", fmt.Errorf("Cisco linecard index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("0/%d/CPU0", index), nil
}

// ControllerCard is a implementation of namer.ControllerCard.
func (n *Namer) ControllerCard(index int) (string, error) {
	const maxIndex = 1
	if index > maxIndex {
		return "", fmt.Errorf("Cisco controller card index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("0/RP%d/CPU0", index), nil
}

// Fabric is a implementation of namer.Fabric.
func (n *Namer) Fabric(index int) (string, error) {
	const maxIndex = 7
	if index > maxIndex {
		return "", fmt.Errorf("Cisco fabric index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("0/FC%d", index), nil
}
