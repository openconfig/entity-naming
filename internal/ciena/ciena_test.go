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

package ciena

import (
	"strings"
	"testing"

	"github.com/openconfig/entity-naming/internal/namer"
)

var cn = new(Namer)

func TestLoopbackInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index uint
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "loop0",
	}, {
		desc:  "max",
		index: 509,
		want:  "loop509",
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

	t.Run("over max", func(t *testing.T) {
		_, err := cn.LoopbackInterface(510)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("LoopbackInterface(256) got error %v, want substring %q", err, wantErr)
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
		want:  "agg1",
	}, {
		desc:  "max",
		index: 255,
		want:  "agg255",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := cn.AggregateInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateInteface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := cn.AggregateInterface(256)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateInterface(65535) got error %v, want substring %q", err, wantErr)
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
		want:  "agg1",
	}, {
		desc:  "max",
		index: 255,
		want:  "agg255",
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
		_, err := cn.AggregateMemberInterface(256)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateMemberInterface(256) got error %v, want substring %q", err, wantErr)
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
		desc: "unchannelizable",
		pp: &namer.PortParams{
			SlotIndex: uintPtr(4),
			PortIndex: 3,
		},
		want: "1/4/3",
	}, {
		desc: "channelized",
		pp: &namer.PortParams{
			SlotIndex:     uintPtr(4),
			PortIndex:     3,
			ChannelIndex:  uintPtr(1),
			Channelizable: true,
		},
		want: "1/4/3",
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
		desc          string
		hardwareModel string
		index         uint
		want          string
		wantErr       bool
	}{{
		desc:          "WR13 default - min slot 1",
		hardwareModel: "",
		index:         1,
		want:          "ib-1/1",
	}, {
		desc:          "WR13 default - slot 6",
		hardwareModel: "",
		index:         6,
		want:          "ib-1/6",
	}, {
		desc:          "WR13 default - slot 10",
		hardwareModel: "",
		index:         10,
		want:          "ib-1/10",
	}, {
		desc:          "WR13 default - slot 11",
		hardwareModel: "",
		index:         11,
		want:          "ib-1/11",
	}, {
		desc:          "WR13 default - second chassis",
		hardwareModel: "",
		index:         17,
		want:          "ib-2/1",
	}, {
		desc:          "WR13 explicit - slot 2",
		hardwareModel: "WR13",
		index:         2,
		want:          "ib-1/2",
	}, {
		desc:          "WR7 - slot 4",
		hardwareModel: "WR7",
		index:         4,
		want:          "ib-1/4",
	}, {
		desc:          "WR7 - slot 5",
		hardwareModel: "WR7",
		index:         5,
		want:          "ib-1/5",
	}, {
		desc:          "WR7 - slot 6",
		hardwareModel: "WR7",
		index:         6,
		want:          "ib-1/6",
	}, {
		desc:          "WR7 - slot 7",
		hardwareModel: "WR7",
		index:         7,
		want:          "ib-1/7",
	}, {
		desc:          "WR2 - slot 4",
		hardwareModel: "WR2",
		index:         4,
		want:          "ib-1/4",
	}, {
		desc:          "WR2 - slot 5",
		hardwareModel: "WR2",
		index:         5,
		want:          "ib-1/5",
	}, {
		desc:          "WR13 - invalid slot 8",
		hardwareModel: "WR13",
		index:         8,
		wantErr:       true,
	}, {
		desc:          "WR7 - invalid slot 3",
		hardwareModel: "WR7",
		index:         3,
		wantErr:       true,
	}, {
		desc:          "WR7 - invalid slot 8",
		hardwareModel: "WR7",
		index:         8,
		wantErr:       true,
	}, {
		desc:          "WR2 - invalid slot 3",
		hardwareModel: "WR2",
		index:         3,
		wantErr:       true,
	}, {
		desc:          "WR2 - invalid slot 6",
		hardwareModel: "WR2",
		index:         6,
		wantErr:       true,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			namer := &Namer{HardwareModel: test.hardwareModel}
			got, err := namer.Linecard(test.index)
			if test.wantErr {
				if err == nil {
					t.Fatalf("Linecard(%v) expected error but got none", test.index)
				}
				return
			}
			if err != nil {
				t.Fatalf("Linecard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Linecard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("default hardware model - invalid slot", func(t *testing.T) {
		_, err := cn.Linecard(8)
		if wantErr := "must be in"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Linecard(8) got error %v, want substring %q", err, wantErr)
		}
	})

	t.Run("unsupported hardware model", func(t *testing.T) {
		namer := &Namer{HardwareModel: "WR99"}
		_, err := namer.Linecard(1)
		if wantErr := "unsupported hardware model"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Linecard with unsupported hardware model got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestControllerCard(t *testing.T) {
	tests := []struct {
		desc          string
		hardwareModel string
		index         uint
		want          string
		wantErr       bool
	}{{
		desc:          "WR13 default - min",
		hardwareModel: "",
		index:         7,
		want:          "ctm-1/7",
	}, {
		desc:          "WR13 default - max",
		hardwareModel: "",
		index:         23,
		want:          "ctm-2/7",
	}, {
		desc:          "WR13 explicit - slot 7",
		hardwareModel: "WR13",
		index:         7,
		want:          "ctm-1/7",
	}, {
		desc:          "WR13 explicit - slot 8",
		hardwareModel: "WR13",
		index:         8,
		want:          "ctm-1/8",
	}, {
		desc:          "WR7 - slot 2",
		hardwareModel: "WR7",
		index:         2,
		want:          "ctm-1/2",
	}, {
		desc:          "WR7 - slot 3",
		hardwareModel: "WR7",
		index:         3,
		want:          "ctm-1/3",
	}, {
		desc:          "WR2 - slot 2",
		hardwareModel: "WR2",
		index:         2,
		want:          "ctm-1/2",
	}, {
		desc:          "WR2 - slot 3",
		hardwareModel: "WR2",
		index:         3,
		want:          "ctm-1/3",
	}, {
		desc:          "WR13 - invalid slot 3",
		hardwareModel: "WR13",
		index:         3,
		wantErr:       true,
	}, {
		desc:          "WR7 - invalid slot 7",
		hardwareModel: "WR7",
		index:         7,
		wantErr:       true,
	}, {
		desc:          "WR2 - invalid slot 8",
		hardwareModel: "WR2",
		index:         8,
		wantErr:       true,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			namer := &Namer{HardwareModel: test.hardwareModel}
			got, err := namer.ControllerCard(test.index)
			if test.wantErr {
				if err == nil {
					t.Fatalf("ControllerCard(%v) expected error but got none", test.index)
				}
				return
			}
			if err != nil {
				t.Fatalf("ControllerCard(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("ControllerCard(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("default hardware model - invalid slot", func(t *testing.T) {
		_, err := cn.ControllerCard(3)
		if wantErr := "must be in"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("ControllerCard(3) got error %v, want substring %q", err, wantErr)
		}
	})

	t.Run("unsupported hardware model", func(t *testing.T) {
		namer := &Namer{HardwareModel: "WR99"}
		_, err := namer.ControllerCard(7)
		if wantErr := "unsupported hardware model"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("ControllerCard with unsupported hardware model got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestFabric(t *testing.T) {
	tests := []struct {
		desc          string
		hardwareModel string
		index         uint
		want          string
		wantErr       bool
	}{{
		desc:          "WR13 default - min slot 12",
		hardwareModel: "",
		index:         12,
		want:          "fb-1/12",
	}, {
		desc:          "WR13 default - max slot 16",
		hardwareModel: "",
		index:         16,
		want:          "fb-1/16",
	}, {
		desc:          "WR13 default - second chassis",
		hardwareModel: "",
		index:         32,
		want:          "fb-2/16",
	}, {
		desc:          "WR13 explicit - slot 13",
		hardwareModel: "WR13",
		index:         13,
		want:          "fb-1/13",
	}, {
		desc:          "WR13 explicit - slot 15",
		hardwareModel: "WR13",
		index:         15,
		want:          "fb-1/15",
	}, {
		desc:          "WR7 - slot 8",
		hardwareModel: "WR7",
		index:         8,
		want:          "fb-1/8",
	}, {
		desc:          "WR7 - slot 9",
		hardwareModel: "WR7",
		index:         9,
		want:          "fb-1/9",
	}, {
		desc:          "WR7 - slot 10",
		hardwareModel: "WR7",
		index:         10,
		want:          "fb-1/10",
	}, {
		desc:          "WR2 - not supported",
		hardwareModel: "WR2",
		index:         1,
		wantErr:       true,
	}, {
		desc:          "WR13 - invalid slot 8",
		hardwareModel: "WR13",
		index:         8,
		wantErr:       true,
	}, {
		desc:          "WR13 - invalid slot 11",
		hardwareModel: "WR13",
		index:         11,
		wantErr:       true,
	}, {
		desc:          "WR7 - invalid slot 7",
		hardwareModel: "WR7",
		index:         7,
		wantErr:       true,
	}, {
		desc:          "WR7 - invalid slot 11",
		hardwareModel: "WR7",
		index:         11,
		wantErr:       true,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			namer := &Namer{HardwareModel: test.hardwareModel}
			got, err := namer.Fabric(test.index)
			if test.wantErr {
				if err == nil {
					t.Fatalf("Fabric(%v) expected error but got none", test.index)
				}
				return
			}
			if err != nil {
				t.Fatalf("Fabric(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("Fabric(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("default hardware model - invalid slot", func(t *testing.T) {
		_, err := cn.Fabric(8)
		if wantErr := "must be in"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Fabric(8) got error %v, want substring %q", err, wantErr)
		}
	})

	t.Run("WR2 not supported", func(t *testing.T) {
		namer := &Namer{HardwareModel: "WR2"}
		_, err := namer.Fabric(12)
		if wantErr := "not supported for WR2"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Fabric with WR2 got error %v, want substring %q", err, wantErr)
		}
	})

	t.Run("unsupported hardware model", func(t *testing.T) {
		namer := &Namer{HardwareModel: "WR99"}
		_, err := namer.Fabric(12)
		if wantErr := "unsupported hardware model"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("Fabric with unsupported hardware model got error %v, want substring %q", err, wantErr)
		}
	})
}
