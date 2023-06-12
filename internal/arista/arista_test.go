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

package arista

import (
	"strings"
	"testing"

	"github.com/openconfig/entity-naming/internal/namer"
)

var an = new(Namer)

func TestLoopbackInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "Loopback0",
	}, {
		desc:  "max",
		index: 1000,
		want:  "Loopback1000",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := an.LoopbackInterface(test.index)
			if err != nil {
				t.Fatalf("LoopbackInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("LoopbackInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := an.LoopbackInterface(1001)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("LoopbackInterface(1001) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestAggregateInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "Port-Channel1",
	}, {
		desc:  "max",
		index: 999998,
		want:  "Port-Channel999999",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := an.AggregateInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateInteface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := an.AggregateInterface(999999)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateInterface(999999) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestAggregateMemberInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "Port-Channel1",
	}, {
		desc:  "max",
		index: 999998,
		want:  "Port-Channel999999",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := an.AggregateMemberInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateMemberInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateMemberInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := an.AggregateMemberInterface(999999)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateMemberInterface(999999) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestPort(t *testing.T) {
	intPtr := func(i int) *int { return &i }

	tests := []struct {
		desc string
		pp   *namer.PortParams
		want string
	}{{
		desc: "unchannelizable",
		pp: &namer.PortParams{
			SlotIndex: intPtr(1),
			PortIndex: 3,
		},
		want: "Ethernet4/3",
	}, {
		desc: "channelizable",
		pp: &namer.PortParams{
			SlotIndex:     intPtr(1),
			PortIndex:     3,
			Channelizable: true,
		},
		want: "Ethernet4/3/1",
	}, {
		desc: "channelized",
		pp: &namer.PortParams{
			SlotIndex:     intPtr(1),
			PortIndex:     3,
			ChannelIndex:  intPtr(4),
			Channelizable: true,
		},
		want: "Ethernet4/3/4",
	}, {
		desc: "fixed form factor - unchannelizable",
		pp: &namer.PortParams{
			PortIndex: 3,
		},
		want: "Ethernet3",
	}, {
		desc: "fixed form factor - channelizable",
		pp: &namer.PortParams{
			PortIndex:     3,
			Channelizable: true,
		},
		want: "Ethernet3/1",
	}, {
		desc: "fixed form factor - channelized",
		pp: &namer.PortParams{
			PortIndex:     3,
			ChannelIndex:  intPtr(4),
			Channelizable: true,
		},
		want: "Ethernet3/4",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := an.Port(test.pp)
			if err != nil {
				t.Fatalf("Port(%v) got error: %v", test.pp, err)
			}
			if got != test.want {
				t.Errorf("Port(%v) got %q, want %q", test.pp, got, test.want)
			}
		})
	}
}

func TestLinecard(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "Linecard3",
	}, {
		desc:  "max",
		index: 7,
		want:  "Linecard10",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := an.Linecard(test.index)
			if err != nil {
				t.Fatalf("Linecard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Linecard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := an.Linecard(8)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Linecard(8) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestControllerCard(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "Supervisor1",
	}, {
		desc:  "max",
		index: 1,
		want:  "Supervisor2",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := an.ControllerCard(test.index)
			if err != nil {
				t.Fatalf("ControllerCard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("ControllerCard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := an.ControllerCard(3)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("ControllerCard(3) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestFabric(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "Fabric1",
	}, {
		desc:  "max",
		index: 5,
		want:  "Fabric6",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := an.Fabric(test.index)
			if err != nil {
				t.Fatalf("Fabric(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Fabric(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := an.Fabric(6)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Fabric(6) got error %v, want substring %q", err, wantErr)
		}
	})
}
