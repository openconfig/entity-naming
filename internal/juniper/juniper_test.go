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

package juniper

import (
	"strings"
	"testing"

	"github.com/openconfig/entity-naming/internal/namer"
)

var jn = new(Namer)

func TestLoopbackInterface(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		got, err := jn.LoopbackInterface(0)
		if err != nil {
			t.Fatalf("LoopbackInterface(0) got error: %v", err)
		}
		if want := "lo0"; got != want {
			t.Errorf("LoopbackInterface(0) got %q, want %q", got, want)
		}
	})

	t.Run("nonzero", func(t *testing.T) {
		_, err := jn.LoopbackInterface(1)
		if wantErr := "zero"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("LoopbackInterface(1) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestAggregateInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index uint
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "ae0",
	}, {
		desc:  "max",
		index: 1151,
		want:  "ae1151",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.AggregateInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := jn.AggregateInterface(1152)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateInterface(1152) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestAggregateMemberInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index uint
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "ae0.0",
	}, {
		desc:  "max",
		index: 1151,
		want:  "ae1151.0",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.AggregateMemberInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateMemberInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateMemberInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := jn.AggregateMemberInterface(1152)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateMemberInterface(1152) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestPort(t *testing.T) {
	uintPtr := func(i uint) *uint { return &i }

	tests := []struct {
		desc string
		pp   *namer.PortParams
		want string
	}{{
		desc: "channelizable",
		pp: &namer.PortParams{
			SlotIndex:     uintPtr(1),
			PICIndex:      2,
			PortIndex:     3,
			Channelizable: true,
		},
		want: "et-1/0/3",
	}, {
		desc: "channelized",
		pp: &namer.PortParams{
			SlotIndex:     uintPtr(1),
			PICIndex:      2,
			PortIndex:     3,
			ChannelIndex:  uintPtr(4),
			Channelizable: true,
		},
		want: "et-1/0/3:4",
	}, {
		desc: "fixed form factor - channelizable",
		pp: &namer.PortParams{
			PICIndex:      2,
			PortIndex:     3,
			Channelizable: true,
		},
		want: "et-0/2/3",
	}, {
		desc: "fixed form factor - channelized",
		pp: &namer.PortParams{
			PICIndex:      2,
			PortIndex:     3,
			ChannelIndex:  uintPtr(4),
			Channelizable: true,
		},
		want: "et-0/2/3:4",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.Port(test.pp)
			if err != nil {
				t.Fatalf("Port(%v) got error: %v", test.pp, err)
			}
			if got != test.want {
				t.Errorf("Port(%v) got %q, want %q", test.pp, got, test.want)
			}
		})
	}

	t.Run("unchannelizable", func(t *testing.T) {
		pp := &namer.PortParams{
			SlotIndex: uintPtr(1),
			PICIndex:  2,
			PortIndex: 3,
		}
		if _, err := jn.Port(pp); err == nil || !strings.Contains(err.Error(), "unchannelizable") {
			t.Fatalf("Port(%v) got unexpected error %v, want substring 'unchannelizable'", pp, err)
		}

	})
}

func TestLinecard(t *testing.T) {
	tests := []struct {
		desc  string
		index uint
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "FPC0",
	}, {
		desc:  "max",
		index: 7,
		want:  "FPC7",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.Linecard(test.index)
			if err != nil {
				t.Fatalf("Linecard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Linecard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := jn.Linecard(8)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Linecard(8) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestControllerCard(t *testing.T) {
	tests := []struct {
		desc  string
		index uint
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "RE0",
	}, {
		desc:  "max",
		index: 1,
		want:  "RE1",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.ControllerCard(test.index)
			if err != nil {
				t.Fatalf("ControllerCard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("ControllerCard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := jn.ControllerCard(3)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("ControllerCard(3) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestFabric(t *testing.T) {
	tests := []struct {
		desc  string
		index uint
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "SIB0",
	}, {
		desc:  "max",
		index: 5,
		want:  "SIB5",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.Fabric(test.index)
			if err != nil {
				t.Fatalf("Fabric(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Fabric(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := jn.Fabric(6)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Fabric(6) got error %v, want substring %q", err, wantErr)
		}
	})
}
