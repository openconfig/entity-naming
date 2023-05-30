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
type Namer struct{}

// LoopbackInterface is a Cisco implementation of namer.LoopbackInterface.
func (n *Namer) LoopbackInterface(index int) (string, error) {
	return fmt.Sprintf("Loopback%d", index), nil
}

// AggregatePort is a Cisco implementation of namer.AggregatePort.
func (n *Namer) AggregatePort(index int) (string, error) {
	index++ // Cisco uses 1-based indices for aggregate ports.
	const maxIndex = 65535
	if index > maxIndex {
		return "", fmt.Errorf("Cisco aggregate index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("Bundle-Ether%d", index), nil
}

// AggregateInterface is a Cisco implementation of namer.AggregateInterface.
func (n *Namer) AggregateInterface(index int) (string, error) {
	return n.AggregatePort(index)
}
