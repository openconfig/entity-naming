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

// Package juniper provides Juniper-specific platform information.
package juniper

import (
	"fmt"

	"github.com/openconfig/entity-naming/internal/namer"
)

var _ namer.Namer = (*Namer)(nil)

// Namer is an Juniper implementation of the namer interface.
type Namer struct{}

// LoopbackInterface is a Juniper implementation of namer.LoopbackInterface.
func (n *Namer) LoopbackInterface(index int) (string, error) {
	const maxIndex = 16000
	if index > maxIndex {
		return "", fmt.Errorf("Juniper loopback index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("lo0.%d", index), nil
}
