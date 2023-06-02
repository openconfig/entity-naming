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

	"github.com/openconfig/entity-naming/internal/namer"
)

var _ namer.Namer = (*Namer)(nil)

// Namer is a Juniper implementation of the Namer interface.
type Namer struct {
	HardwareModel string
}

// LoopbackInterface is a Juniper implementation of namer.LoopbackInterface.
func (n *Namer) LoopbackInterface(index int) (string, error) {
	const maxIndex = 16000
	if index > maxIndex {
		return "", fmt.Errorf("Juniper loopback index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("lo0.%d", index), nil
}

// AggregateInterface is a Juniper implementation of namer.AggregateInterface.
func (n *Namer) AggregateInterface(index int) (string, error) {
	const maxIndex = 1151
	if index > maxIndex {
		return "", fmt.Errorf("Juniper aggregate index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("ae%d", index), nil
}

// AggregateMemberInterface is a Juniper implementation of namer.AggregateMemberInterface.
func (n *Namer) AggregateMemberInterface(index int) (string, error) {
	name, err := n.AggregateInterface(index)
	if err != nil {
		return "", err
	}
	return name + ".0", nil
}

// Linecard is a Juniper implementation of namer.Linecard.
func (n *Namer) Linecard(index int) (string, error) {
	const maxIndex = 7
	if index > maxIndex {
		return "", fmt.Errorf("Juniper linecard index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("FPC%d", index), nil
}

// ControllerCard is a Juniper implementation of namer.ControllerCard.
func (n *Namer) ControllerCard(index int) (string, error) {
	const maxIndex = 1
	if index > maxIndex {
		return "", fmt.Errorf("Juniper controller card index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("RE%d", index), nil
}

// Fabric is a Juniper implementation of namer.Fabric.
func (n *Namer) Fabric(index int) (string, error) {
	const maxIndex = 5
	if index > maxIndex {
		return "", fmt.Errorf("Juniper fabric index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("SIB%d", index), nil
}
