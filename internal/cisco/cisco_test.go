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

package cisco

import (
	"strings"
	"testing"

	"github.com/openconfig/entity-naming/internal/namer"
	"github.com/openconfig/entity-naming/oc"
)

var cn = new(Namer)

func TestLoopbackInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "zero",
		index: 0,
		want:  "Loopback0",
	}, {
		desc:  "nonzero",
		index: 1000,
		want:  "Loopback1000",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := cn.LoopbackInterface(test.index)
			if err != nil {
				t.Fatalf("LoopbackInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("LoopbackInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}
}

func TestAggregateInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "Bundle-Ether1",
	}, {
		desc:  "max",
		index: 65534,
		want:  "Bundle-Ether65535",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := cn.AggregateInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := cn.AggregateInterface(65535)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateInterface(65535) got error %v, want substring %q", err, wantErr)
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
		want:  "Bundle-Ether1",
	}, {
		desc:  "max",
		index: 65534,
		want:  "Bundle-Ether65535",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := cn.AggregateMemberInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateMemberInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateMemberInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := cn.AggregateMemberInterface(65535)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateMemberInterface(65535) got error %v, want substring %q", err, wantErr)
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
			Speed:     oc.IfEthernet_ETHERNET_SPEED_SPEED_10GB,
		},
		want: "TenGigE0/1/0/3",
	}, {
		desc: "channelizable",
		pp: &namer.PortParams{
			SlotIndex:     intPtr(1),
			PortIndex:     3,
			Speed:         oc.IfEthernet_ETHERNET_SPEED_SPEED_100GB,
			Channelizable: true,
		},
		want: "HundredGigE0/1/0/3",
	}, {
		desc: "channelized",
		pp: &namer.PortParams{
			SlotIndex:    intPtr(1),
			PortIndex:    3,
			ChannelIndex: intPtr(4),
			Speed:        oc.IfEthernet_ETHERNET_SPEED_SPEED_400GB,
		},
		want: "FourHundredGigE0/1/0/3/4",
	}, {
		desc: "fixed form factor - unchannelizable",
		pp: &namer.PortParams{
			PortIndex: 3,
			Speed:     oc.IfEthernet_ETHERNET_SPEED_SPEED_10GB,
		},
		want: "TenGigE0/0/0/3",
	}, {
		desc: "fixed form factor - channelizable",
		pp: &namer.PortParams{
			PortIndex:     3,
			Speed:         oc.IfEthernet_ETHERNET_SPEED_SPEED_100GB,
			Channelizable: true,
		},
		want: "HundredGigE0/0/0/3",
	}, {
		desc: "channelized fixed form factor",
		pp: &namer.PortParams{
			PortIndex:     3,
			ChannelIndex:  intPtr(4),
			Speed:         oc.IfEthernet_ETHERNET_SPEED_SPEED_400GB,
			Channelizable: true,
		},
		want: "FourHundredGigE0/0/0/3/4",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := cn.Port(test.pp)
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
		want:  "0/0/CPU0",
	}, {
		desc:  "max",
		index: 7,
		want:  "0/7/CPU0",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := cn.Linecard(test.index)
			if err != nil {
				t.Fatalf("Linecard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Linecard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := cn.Linecard(8)
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
		want:  "0/RP0/CPU0",
	}, {
		desc:  "max",
		index: 1,
		want:  "0/RP1/CPU0",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := cn.ControllerCard(test.index)
			if err != nil {
				t.Fatalf("ControllerCard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("ControllerCard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := cn.ControllerCard(3)
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
		want:  "0/FC0",
	}, {
		desc:  "max",
		index: 7,
		want:  "0/FC7",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := cn.Fabric(test.index)
			if err != nil {
				t.Fatalf("Fabric(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Fabric(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := cn.Fabric(8)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Fabric(8) got error %v, want substring %q", err, wantErr)
		}
	})
}
