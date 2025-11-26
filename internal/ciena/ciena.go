// Copyright 2025 Ciena Corp.
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

// Package Ciena provides an Ciena naming implementation.
package ciena

import (
	"fmt"
	"strings"

	"github.com/openconfig/entity-naming/internal/namer"
)

var _ namer.Namer = (*Namer)(nil)

const (
	hardwareModelWR13 = "WR13"
	hardwareModelWR7  = "WR7"
	hardwareModelWR2  = "WR2"
)

// Namer is an Ciena implementation of the Namer interface.
type Namer struct {
	HardwareModel string // WR13 (default), WR7, or WR2
}

// LoopbackInterface is an implementation of namer.LoopbackInterface.
func (n *Namer) LoopbackInterface(index uint) (string, error) {
	const maxIndex = 509
	if index > maxIndex {
		return "", fmt.Errorf("ciena loopback index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("loop%d", index), nil
}

// AggregateInterface is an implementation of namer.AggregateInterface.
func (n *Namer) AggregateInterface(index uint) (string, error) {
	const maxIndex = 255
	if index > maxIndex {
		return "", fmt.Errorf("ciena aggregate index cannot exceed %d, got %d", maxIndex, index)
	}
	return fmt.Sprintf("agg%d", index+1), nil
}

// AggregateMemberInterface is an implementation of namer.AggregateMemberInterface.
func (n *Namer) AggregateMemberInterface(index uint) (string, error) {
	return n.AggregateInterface(index)
}

// calculateSlotIndices calculates the hardware and slot indices from a linear index.
// hIndex represents the hardware/chassis index, sIndex represents the slot index.
func calculateSlotIndices(index uint) (hIndex, sIndex uint) {
	hIndex = ((index - 1) / 16) + 1
	sIndex = ((index - 1) % 16) + 1
	return hIndex, sIndex
}

// Linecard is an implementation of namer.Linecard.
func (n *Namer) Linecard(index uint) (string, error) {
	hIndex, sIndex := calculateSlotIndices(index)

	// Default to WR13 if HardwareModel is not set
	hardwareModel := n.HardwareModel
	if hardwareModel == "" {
		hardwareModel = hardwareModelWR13
	}

	switch hardwareModel {
	case hardwareModelWR2:
		// For WR2, linecards are at slots 4 and 5
		switch sIndex {
		case 4, 5:
			return fmt.Sprintf("ib-%d/%d", hIndex, sIndex), nil
		default:
			return "", fmt.Errorf("ciena linecard slot index for WR2 must be in [4,5], got %d", sIndex)
		}
	case hardwareModelWR7:
		// For WR7, linecards are at slots 4, 5, 6, and 7
		switch sIndex {
		case 4, 5, 6, 7:
			return fmt.Sprintf("ib-%d/%d", hIndex, sIndex), nil
		default:
			return "", fmt.Errorf("ciena linecard slot index for WR7 must be in [4,5,6,7], got %d", sIndex)
		}
	case hardwareModelWR13:
		// For WR13, linecards are at slots 1-6 and 10-11
		switch sIndex {
		case 1, 2, 3, 4, 5, 6, 10, 11:
			return fmt.Sprintf("ib-%d/%d", hIndex, sIndex), nil
		default:
			return "", fmt.Errorf("ciena linecard slot index for WR13 must be in [1-6,10,11], got %d", sIndex)
		}
	default:
		return "", fmt.Errorf("unsupported hardware model: %s (supported: WR13, WR7, WR2)", hardwareModel)
	}
}

// ControllerCard is an implementation of namer.ControllerCard.
func (n *Namer) ControllerCard(index uint) (string, error) {
	hIndex, sIndex := calculateSlotIndices(index)

	// Default to WR13 if HardwareModel is not set
	hardwareModel := n.HardwareModel
	if hardwareModel == "" {
		hardwareModel = hardwareModelWR13
	}

	switch hardwareModel {
	case hardwareModelWR7, hardwareModelWR2:
		// For WR7 and WR2, controller cards are at slots 2 and 3
		switch sIndex {
		case 2, 3:
			return fmt.Sprintf("ctm-%d/%d", hIndex, sIndex), nil
		default:
			return "", fmt.Errorf("ciena Controller Card slot index for %s must be in [2,3], got %d", hardwareModel, sIndex)
		}
	case hardwareModelWR13:
		// For WR13, controller cards are at slots 7 and 8
		switch sIndex {
		case 7, 8:
			return fmt.Sprintf("ctm-%d/%d", hIndex, sIndex), nil
		default:
			return "", fmt.Errorf("ciena Controller Card slot index for WR13 must be in [7,8], got %d", sIndex)
		}
	default:
		return "", fmt.Errorf("unsupported hardware model: %s (supported: WR13, WR7, WR2)", hardwareModel)
	}
}

// Fabric is an implementation of namer.Fabric.
func (n *Namer) Fabric(index uint) (string, error) {
	hIndex, sIndex := calculateSlotIndices(index)

	// Default to WR13 if HardwareModel is not set
	hardwareModel := n.HardwareModel
	if hardwareModel == "" {
		hardwareModel = hardwareModelWR13
	}

	switch hardwareModel {
	case hardwareModelWR2:
		// For WR2, fabric cards are not supported
		return "", fmt.Errorf("ciena Fabric is not supported for WR2")
	case hardwareModelWR7:
		// For WR7, fabric cards are at slots 8, 9, and 10
		switch sIndex {
		case 8, 9, 10:
			return fmt.Sprintf("fb-%d/%d", hIndex, sIndex), nil
		default:
			return "", fmt.Errorf("ciena Fabric slot index for WR7 must be in [8,9,10], got %d", sIndex)
		}
	case hardwareModelWR13:
		// For WR13, fabric cards are at slots 12-16
		switch sIndex {
		case 12, 13, 14, 15, 16:
			return fmt.Sprintf("fb-%d/%d", hIndex, sIndex), nil
		default:
			return "", fmt.Errorf("ciena Fabric slot index for WR13 must be in [12-16], got %d", sIndex)
		}
	default:
		return "", fmt.Errorf("unsupported hardware model: %s (supported: WR13, WR7, WR2)", hardwareModel)
	}
}

// Port is an implementation of namer.Port.
func (n *Namer) Port(pp *namer.PortParams) (string, error) {
	var nameBuilder strings.Builder
	if pp.ChannelIndex == nil {
		nameBuilder.WriteString("1/")
	} else {
		nameBuilder.WriteString(fmt.Sprintf("%d/", *pp.ChannelIndex))
	}
	if pp.SlotIndex != nil {
		nameBuilder.WriteString(fmt.Sprintf("%d/", *pp.SlotIndex))
	}
	nameBuilder.WriteString(fmt.Sprintf("%d", pp.PortIndex))

	return nameBuilder.String(), nil
}

// IsFixedFormFactor is an implementation of namer.IsFixedFormFactor.
func (n *Namer) IsFixedFormFactor() bool {
	// TODO(Ciena): Fill in this implementation.
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
