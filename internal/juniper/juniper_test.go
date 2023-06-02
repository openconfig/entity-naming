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
)

var jn = new(Namer)

func TestLoopbackInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "lo0.0",
	}, {
		desc:  "max",
		index: 16000,
		want:  "lo0.16000",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.LoopbackInterface(test.index)
			if err != nil {
				t.Fatalf("LoopbackInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("LoopbackInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := jn.LoopbackInterface(16001)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("LoopbackInterface(16001) got error %v, want substring %q", err, wantErr)
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
		index int
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

func TestLinecard(t *testing.T) {
	tests := []struct {
		desc  string
		index int
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
		index int
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
		index int
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
