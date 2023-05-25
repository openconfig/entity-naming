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

func TestAggregatePort(t *testing.T) {
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
			got, err := cn.AggregatePort(test.index)
			if err != nil {
				t.Fatalf("AggregatePort(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregatePort(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := cn.AggregatePort(65535)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregatePort(65535) got error %v, want substring %q", err, wantErr)
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
