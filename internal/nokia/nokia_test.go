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

package nokia

import (
	"strings"
	"testing"

	"github.com/openconfig/entity-naming/internal/namer"
)

var nn = new(Namer)

func TestLoopbackInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "lo0",
	}, {
		desc:  "max",
		index: 255,
		want:  "lo255",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := nn.LoopbackInterface(test.index)
			if err != nil {
				t.Fatalf("LoopbackInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("LoopbackInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := nn.LoopbackInterface(256)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("LoopbackInterface(256) got error %v, want substring %q", err, wantErr)
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
		want:  "lag1",
	}, {
		desc:  "max",
		index: 127,
		want:  "lag128",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := nn.AggregateInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := nn.AggregateInterface(128)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateInterface(128) got error %v, want substring %q", err, wantErr)
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
		want:  "lag1.0",
	}, {
		desc:  "max",
		index: 127,
		want:  "lag128.0",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := nn.AggregateMemberInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateMemberInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateMemberInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := nn.AggregateMemberInterface(128)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateMemberInterface(128) got error %v, want substring %q", err, wantErr)
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
		desc: "standard",
		pp: &namer.PortParams{
			SlotIndex: intPtr(1),
			PortIndex: 3,
		},
		want: "et-2/4",
	}, {
		desc: "channelized",
		pp: &namer.PortParams{
			SlotIndex:    intPtr(1),
			PortIndex:    3,
			ChannelIndex: intPtr(4),
		},
		want: "et-2/4/5",
	}, {
		desc: "fixed form factor",
		pp: &namer.PortParams{
			PortIndex: 3,
		},
		want: "et-1/4",
	}, {
		desc: "channelized fixed form factor",
		pp: &namer.PortParams{
			PortIndex:    3,
			ChannelIndex: intPtr(4),
		},
		want: "et-1/4/5",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := nn.Port(test.pp)
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
		want:  "Linecard1",
	}, {
		desc:  "max",
		index: 7,
		want:  "Linecard8",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := nn.Linecard(test.index)
			if err != nil {
				t.Fatalf("Linecard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Linecard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := nn.Linecard(8)
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
			got, err := nn.ControllerCard(test.index)
			if err != nil {
				t.Fatalf("ControllerCard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("ControllerCard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := nn.ControllerCard(3)
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
		index: 7,
		want:  "Fabric8",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := nn.Fabric(test.index)
			if err != nil {
				t.Fatalf("Fabric(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Fabric(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := nn.Fabric(8)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Fabric(8) got error %v, want substring %q", err, wantErr)
		}
	})
}
